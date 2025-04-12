package reception

import (
	base "avito-pvz/internal/handlers/dto"
	"avito-pvz/internal/metrics"
	"avito-pvz/internal/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Create handler checks if there are any open receptions and if not, creates a new one
func (h *Handler) Create(ctx echo.Context) error {
	var req base.ReceptionRequest
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

	reception := models.Reception{
		ID:       uuid.New(),
		DateTime: time.Now(),
		PvzID:    req.PvzID,
		Status:   string(base.InProgressStatus),
	}

	if userRole.IsEmployee() {
		lastReception, err := h.ReceptionUU.GetLastReception(ctx.Request().Context(), reception.PvzID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
				Message: "failed to fetch last reception",
			})
		}
		if base.ReceptionStatus(lastReception.Status) == base.InProgressStatus {
			return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
				Message: "unable to start another reception",
			})
		}
		err = h.ReceptionUU.StartReception(ctx.Request().Context(), reception)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
				Message: "failed to add product",
			})
		}
	} else {
		return ctx.JSON(http.StatusForbidden, base.ErrorResponse{
			Message: "access denied: insufficient permissions",
		})
	}

	metrics.IncrementCreatedReceptions()
	return ctx.JSON(http.StatusCreated, base.Reception{
		ID:       reception.ID,
		DateTime: reception.DateTime,
		PvzID:    req.PvzID,
		Status:   base.ReceptionStatus(reception.Status),
	})

}
