package auth

import (
	"avito-pvz/internal/auth"
	base "avito-pvz/internal/handlers"
	"avito-pvz/internal/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Handler struct {
	UserUU *usecase.UserUseCase
	logger *slog.Logger
}

func NewAuthHandler(uuu *usecase.UserUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		UserUU: uuu,
		logger: logger,
	}
}

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

	token, err := auth.GenerateToken("mock@mail.ru", string(req.Role))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.ErrorResponse{
			Message: "failed to generate token",
		})
	}

	return ctx.JSON(http.StatusOK, base.DummyResponse{
		Token: token,
	})
}
