package auth

import (
	"avito-pvz/internal/auth"
	base "avito-pvz/internal/handlers/dto"
	"avito-pvz/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Login handler used to authenticate user and issue a token
func (h *Handler) Login(ctx echo.Context) error {
	var req base.LoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid request body",
		})
	}

	exists, err := h.UserUU.CheckUserExists(ctx.Request().Context(), req.Email)
	if err != nil || !exists {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "failed to find user",
		})
	}

	var user models.User
	user, err = h.UserUU.GetUserByEmail(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to fetch user",
		})
	}

	if correct := auth.CheckPasswordHash(req.Password, user.Hash); !correct {
		return ctx.JSON(http.StatusUnauthorized, base.ErrorResponse{
			Message: "invalid user password",
		})
	}

	token, err := auth.GenerateToken(user.Email, user.Role)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to generate token",
		})
	}

	return ctx.JSON(http.StatusOK, base.DummyResponse{
		Token: token,
	})
}
