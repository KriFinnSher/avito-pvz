package auth

import (
	"avito-pvz/internal/auth"
	base "avito-pvz/internal/handlers"
	"avito-pvz/internal/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

// Handler structure stands for all authorization handlers, e.g. /dummyLogin, /login and /register
type Handler struct {
	UserUU *usecase.UserUseCase
	logger *slog.Logger
}

// NewAuthHandler simply creates new Handler instance
func NewAuthHandler(uuu *usecase.UserUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		UserUU: uuu,
		logger: logger,
	}
}

// DummyLogin handler receiving role and responding with token (only for tests)
func (h *Handler) DummyLogin(ctx echo.Context) error {
	var req base.DummyRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
			Message: "invalid request body",
		})
	}

	//email, ok := ctx.Get("email").(string)
	//if !ok {
	//	return ctx.JSON(http.StatusBadRequest, base.ErrorResponse{
	//		Message: "email not found in context",
	//	})
	//}

	token, err := auth.GenerateToken("mocks@mail.ru", string(req.Role))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to generate token",
		})
	}

	return ctx.JSON(http.StatusOK, base.DummyResponse{
		Token: token,
	})
}
