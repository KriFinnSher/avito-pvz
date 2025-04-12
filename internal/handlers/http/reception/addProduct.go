package reception

import (
	base "avito-pvz/internal/handlers/dto"
	"avito-pvz/internal/metrics"
	"avito-pvz/internal/models"
	"avito-pvz/internal/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Handler structure stands for all reception manipulating handlers
type Handler struct {
	ReceptionUU *usecase.ReceptionUseCase
	ProductUU   *usecase.ProductUseCase
}

// NewReceptionHandler creates new instance of Handler
func NewReceptionHandler(ruu *usecase.ReceptionUseCase, puu *usecase.ProductUseCase) *Handler {
	return &Handler{
		ReceptionUU: ruu,
		ProductUU:   puu,
	}
}

// AddProduct handler creates new product and adds it to the last open reception of pvz with specific id
func (h *Handler) AddProduct(ctx echo.Context) error {
	var req base.ProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid request body",
		})
	}

	reception, err := h.ReceptionUU.GetLastReception(ctx.Request().Context(), req.PvzID)
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

	if !req.Type.Validate() {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid product type",
		})
	}

	product := models.Product{
		ID:          uuid.New(),
		DateTime:    time.Now(),
		Type:        string(req.Type),
		ReceptionId: reception.ID,
	}

	if userRole.IsEmployee() {
		err = h.ProductUU.AddNew(ctx.Request().Context(), product)
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

	metrics.IncrementAddedProducts()
	return ctx.JSON(http.StatusCreated, base.Product{
		ID:          product.ID,
		DateTime:    product.DateTime,
		Type:        base.ProductType(product.Type),
		ReceptionID: reception.ID,
	})
}
