package pvz

import (
	base "avito-pvz/internal/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// GetAll handler show whole information about pvz and its internals
func (h *Handler) GetAll(ctx echo.Context) error {
	startDateStr := ctx.QueryParam("startDate")
	endDateStr := ctx.QueryParam("endDate")
	pageStr := ctx.QueryParam("page")
	limitStr := ctx.QueryParam("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{Message: "Invalid page parameter"})
		}
	}

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 || limit > 30 {
			return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{Message: "Invalid limit parameter"})
		}
	}

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{Message: "Invalid startDate parameter"})
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{Message: "Invalid endDate parameter"})
		}
	}

	PVZs, err := h.PvzUU.GetAllPVZs(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{Message: "Failed to get PVZs"})
	}

	var filtered []base.PVZResponse

	for _, pvz := range PVZs {
		if (!startDate.IsZero() && pvz.RegistrationDate.Before(startDate)) || (!endDate.IsZero() && pvz.RegistrationDate.After(endDate)) {
			continue
		}

		receptions, err := h.ReceptionUU.GetListForPVZ(ctx.Request().Context(), pvz.ID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{Message: "Failed to get receptions"})
		}

		var receptionWithProducts []base.ReceptionWithProducts

		for _, reception := range receptions {
			products, err := h.ProductUU.GetReceptionList(ctx.Request().Context(), reception.ID)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{Message: "Failed to get products"})
			}

			var dtoProducts []base.Product
			for _, p := range products {
				dtoProducts = append(dtoProducts, base.Product{
					ID:          p.ID,
					DateTime:    p.DateTime,
					Type:        base.ProductType(p.Type),
					ReceptionID: p.ReceptionId,
				})
			}

			receptionWithProducts = append(receptionWithProducts, base.ReceptionWithProducts{
				Reception: base.Reception{
					ID:       reception.ID,
					DateTime: reception.DateTime,
					PvzID:    reception.PvzID,
					Status:   base.ReceptionStatus(reception.Status),
				},
				Products: dtoProducts,
			})
		}

		filtered = append(filtered, base.PVZResponse{
			PVZ: base.PVZ{
				ID:               pvz.ID,
				RegistrationDate: pvz.RegistrationDate,
				City:             base.PvzCity(pvz.City),
			},
			Receptions: receptionWithProducts,
		})
	}

	total := len(filtered)
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	return ctx.JSON(http.StatusOK, filtered[start:end])
}
