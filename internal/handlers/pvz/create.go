package pvz

import (
	base "avito-pvz/internal/handlers"
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

	userRole, ok := ctx.Get("role").(string)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to fetch user role",
		})
	}

	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}
	if req.RegistrationDate.IsZero() {
		req.RegistrationDate = time.Now()
	}

	switch base.UserRole(userRole) {
	case base.ModeratorRole:
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
	case base.EmployeeRole:
		return ctx.JSON(http.StatusForbidden, base.ErrorResponse{
			Message: "access denied: insufficient permissions",
		})
	default:
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "invalid user role",
		})
	}

	return ctx.JSON(http.StatusCreated, base.PVZ{
		ID:               req.ID,
		RegistrationDate: req.RegistrationDate,
		City:             req.City,
	})
}
