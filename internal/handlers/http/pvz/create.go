package pvz

import (
	base "avito-pvz/internal/handlers/dto"
	"avito-pvz/internal/metrics"
	"avito-pvz/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Create handler is used to create new pvz
func (h *Handler) Create(ctx echo.Context) error {
	var req base.PVZ

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid request body",
		})
	}

	userRole, ok := ctx.Get("role").(base.UserRole)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to fetch user role",
		})
	}

	if !req.City.Validate() {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid city for pvz",
		})
	}

	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}
	if req.RegistrationDate.IsZero() {
		req.RegistrationDate = time.Now()
	}

	if userRole.IsModerator() {
		pvz := models.PVZ{
			ID:               req.ID,
			RegistrationDate: req.RegistrationDate,
			City:             string(req.City),
		}
		err := h.PvzUU.CreatePVZ(ctx.Request().Context(), pvz)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
				Message: "failed to create pvz",
			})
		}
	} else {
		return ctx.JSON(http.StatusForbidden, base.ErrorResponse{
			Message: "access denied: insufficient permissions",
		})
	}

	metrics.IncrementCreatedPVZs()
	return ctx.JSON(http.StatusCreated, base.PVZ{
		ID:               req.ID,
		RegistrationDate: req.RegistrationDate,
		City:             req.City,
	})
}
