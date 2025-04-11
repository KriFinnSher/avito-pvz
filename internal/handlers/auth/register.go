package auth

import (
	"avito-pvz/internal/auth"
	base "avito-pvz/internal/handlers"
	"avito-pvz/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Register is used to create new user (if not present with specific email)
func (h *Handler) Register(ctx echo.Context) error {
	var req base.RegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid request body",
		})
	}

	exists, err := h.UserUU.CheckUserExists(ctx.Request().Context(), req.Email)
	if err != nil || exists {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "user already exists",
		})
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to hash password",
		})
	}

	user := models.User{
		ID:    uuid.New(),
		Email: req.Email,
		Role:  string(req.Role),
		Hash:  hash,
	}
	if err = h.UserUU.RegisterUser(ctx.Request().Context(), user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to register user",
		})
	}

	return ctx.JSON(http.StatusCreated, base.User{
		ID:    user.ID,
		Email: req.Email,
		Role:  req.Role,
	})
}
