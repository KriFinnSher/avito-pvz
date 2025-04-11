package usecase

import (
	mockPVZ "avito-pvz/internal/mocks/pvz"
	"avito-pvz/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestCreatePVZ_Success(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newPVZ := models.PVZ{
		ID:               pvzID,
		RegistrationDate: time.Now(),
		City:             "Test City",
	}

	pvzRepo.On("Create", ctx, newPVZ).Return(nil)

	useCase := NewPVZUseCase(pvzRepo, logger)

	err := useCase.CreatePVZ(ctx, newPVZ)

	assert.NoError(t, err)
	pvzRepo.AssertCalled(t, "Create", ctx, newPVZ)
}

func TestCreatePVZ_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newPVZ := models.PVZ{
		ID:               pvzID,
		RegistrationDate: time.Now(),
		City:             "Test City",
	}

	pvzRepo.On("Create", ctx, newPVZ).Return(errors.New("failed to create PVZ"))

	useCase := NewPVZUseCase(pvzRepo, logger)

	err := useCase.CreatePVZ(ctx, newPVZ)

	assert.Error(t, err)
	assert.Equal(t, "failed to create PVZ", err.Error())
	pvzRepo.AssertCalled(t, "Create", ctx, newPVZ)
}

func TestGetAllPVZs_Success(t *testing.T) {
	ctx := context.Background()

	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	expectedPVZs := []models.PVZ{
		{ID: uuid.New(), RegistrationDate: time.Now(), City: "City 1"},
		{ID: uuid.New(), RegistrationDate: time.Now(), City: "City 2"},
	}

	pvzRepo.On("GetAll", ctx).Return(expectedPVZs, nil)

	useCase := NewPVZUseCase(pvzRepo, logger)

	pvzs, err := useCase.GetAllPVZs(ctx)

	assert.NoError(t, err)
	assert.Len(t, pvzs, 2)
	assert.Equal(t, expectedPVZs, pvzs)
	pvzRepo.AssertCalled(t, "GetAll", ctx)
}

func TestGetAllPVZs_Error(t *testing.T) {
	ctx := context.Background()

	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	pvzRepo.On("GetAll", ctx).Return(nil, errors.New("failed to fetch PVZs"))

	useCase := NewPVZUseCase(pvzRepo, logger)

	pvzs, err := useCase.GetAllPVZs(ctx)

	assert.Error(t, err)
	assert.Nil(t, pvzs)
	pvzRepo.AssertCalled(t, "GetAll", ctx)
}
