package pvz

import (
	base "avito-pvz/internal/handlers"
	"avito-pvz/internal/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Handler struct {
	PvzUU       *usecase.PVZUseCase
	ReceptionUU *usecase.ReceptionUseCase
	ProductUU   *usecase.ProductUseCase
	logger      *slog.Logger
}

func NewPvzHandler(puu *usecase.PVZUseCase, ruu *usecase.ReceptionUseCase, pruu *usecase.ProductUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		PvzUU:       puu,
		ReceptionUU: ruu,
		ProductUU:   pruu,
		logger:      logger,
	}
}

func (h *Handler) CloseLast(ctx echo.Context) error {
	pvzIdStr := ctx.Param("pvzId")
	pvzId, err := uuid.Parse(pvzIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid pvzId format",
		})
	}

	reception, err := h.ReceptionUU.GetLastReception(ctx.Request().Context(), pvzId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to get last reception",
		})
	}

	userRole, ok := ctx.Get("role").(string)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to fetch user role",
		})
	}

	switch base.UserRole(userRole) {
	case base.EmployeeRole:
		err = h.ReceptionUU.CloseReception(ctx.Request().Context(), reception.ID, pvzId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
				Message: "failed to close reception",
			})
		}
	case base.ModeratorRole:
		return ctx.JSON(http.StatusForbidden, base.ErrorResponse{
			Message: "access denied: insufficient permissions",
		})
	default:
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "invalid user role",
		})
	}

	return ctx.JSON(http.StatusOK, base.Reception{
		ID:       reception.ID,
		DateTime: reception.DateTime,
		PvzID:    reception.PvzID,
		Status:   base.ReceptionStatus(reception.Status),
	})
}
