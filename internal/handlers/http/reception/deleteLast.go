package reception

import (
	base "avito-pvz/internal/handlers/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// DeleteLast handler removes last (LIFO) product from last open reception if present
func (h *Handler) DeleteLast(ctx echo.Context) error {
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

	userRole, ok := ctx.Get("role").(base.UserRole)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to fetch user role",
		})
	}

	if userRole.IsEmployee() {
		err = h.ProductUU.RemoveFromReception(ctx.Request().Context(), reception.ID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
				Message: "failed to remove product",
			})
		}
	} else {
		return ctx.JSON(http.StatusForbidden, base.ErrorResponse{
			Message: "access denied: insufficient permissions",
		})
	}
	var i any
	return ctx.JSON(http.StatusOK, i)
}
