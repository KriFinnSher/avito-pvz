package main

import (
	"avito-pvz/internal/config"
	db "avito-pvz/internal/db/postgres"
	"avito-pvz/internal/grpc/server"
	"avito-pvz/internal/handlers/auth"
	"avito-pvz/internal/handlers/pvz"
	"avito-pvz/internal/handlers/reception"
	mm "avito-pvz/internal/middleware"
	"avito-pvz/internal/repository/postgres"
	"avito-pvz/internal/usecase"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := config.SetUp()
	if err != nil {
		logger.Error("critical error: failed to setup config", "err", err)
	}
	postgresDB, err := db.InitDB()
	if err != nil {
		logger.Error("critical error: failed to setup db", "err", err)
	}
	err = db.MakeMigrations(true)
	if err != nil {
		logger.Error("critical error: failed to make migrations", "err", err)
	}

	e := echo.New()
	e.Use(middleware.Recover())

	userRepo := postgres.NewUserRepo(postgresDB, logger)
	productRepo := postgres.NewProductRepo(postgresDB, logger)
	pvzRepo := postgres.NewPVZRepo(postgresDB, logger)
	receptionRepo := postgres.NewReceptionRepo(postgresDB, logger)

	userUU := usecase.NewUserUseCase(userRepo, logger)
	productUU := usecase.NewProductUseCase(productRepo, receptionRepo, logger)
	pvzUU := usecase.NewPVZUseCase(pvzRepo, logger)
	receptionUU := usecase.NewReceptionUseCase(receptionRepo, logger)

	authHandler := auth.NewAuthHandler(userUU)
	pvzHandler := pvz.NewPvzHandler(pvzUU, receptionUU, productUU)
	receptionHandler := reception.NewReceptionHandler(receptionUU, productUU)

	server.RunGRPCServer(pvzUU, logger)

	e.POST("/dummyLogin", authHandler.DummyLogin)
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	authGroup := e.Group("", mm.JwtMiddleware)

	authGroup.POST("/pvz", pvzHandler.Create)
	authGroup.GET("/pvz", pvzHandler.GetAll)
	authGroup.POST("/pvz/:pvzId/close_last_reception", pvzHandler.CloseLast)
	authGroup.POST("/pvz/:pvzId/delete_last_product", receptionHandler.DeleteLast)
	authGroup.POST("/receptions", receptionHandler.Create)
	authGroup.POST("/products", receptionHandler.AddProduct)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start server", "error", err)
		}
	}()

	<-stop
	logger.Info("received shutdown signal, starting shutdown...")

	// db.MakeMigrations(false)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Error("failed to gracefully shut down server", "error", err)
	}

	logger.Info("server gracefully stopped")
}
