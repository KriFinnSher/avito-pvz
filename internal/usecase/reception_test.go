package usecase

import (
	"avito-pvz/internal/models"
	mockReception "avito-pvz/internal/repository/mocks/reception"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestStartReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "open",
	}

	receptionRepo.On("Create", ctx, newReception).Return(nil)

	useCase := NewReceptionUseCase(receptionRepo, logger)

	err := useCase.StartReception(ctx, newReception)

	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "Create", ctx, newReception)
}

func TestStartReception_Error(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	newReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "open",
	}

	receptionRepo.On("Create", ctx, newReception).Return(errors.New("failed to start reception"))

	useCase := NewReceptionUseCase(receptionRepo, logger)

	err := useCase.StartReception(ctx, newReception)

	assert.Error(t, err)
	receptionRepo.AssertCalled(t, "Create", ctx, newReception)
}

func TestCloseReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(true)
	receptionRepo.On("CloseLast", ctx, pvzID).Return(nil)

	useCase := NewReceptionUseCase(receptionRepo, logger)

	err := useCase.CloseReception(ctx, receptionID, pvzID)

	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	receptionRepo.AssertCalled(t, "CloseLast", ctx, pvzID)
}

func TestCloseReception_Error_AlreadyClosed(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(false)

	useCase := NewReceptionUseCase(receptionRepo, logger)

	err := useCase.CloseReception(ctx, receptionID, pvzID)

	assert.Error(t, err)
	assert.Equal(t, "reception is already closed", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	receptionRepo.AssertNotCalled(t, "CloseLast", ctx, pvzID)
}

func TestGetLastReception_Success(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()
	receptionID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	expectedReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "closed",
	}

	receptionRepo.On("GetLast", ctx, pvzID).Return(expectedReception, nil)

	useCase := NewReceptionUseCase(receptionRepo, logger)

	reception, err := useCase.GetLastReception(ctx, pvzID)

	assert.NoError(t, err)
	assert.Equal(t, expectedReception, reception)
	receptionRepo.AssertCalled(t, "GetLast", ctx, pvzID)
}

func TestGetLastReception_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("GetLast", ctx, pvzID).Return(models.Reception{}, errors.New("failed to fetch last reception"))

	useCase := NewReceptionUseCase(receptionRepo, logger)

	reception, err := useCase.GetLastReception(ctx, pvzID)

	assert.Error(t, err)
	assert.Equal(t, "failed to fetch last reception", err.Error())
	assert.Equal(t, models.Reception{}, reception)
	receptionRepo.AssertCalled(t, "GetLast", ctx, pvzID)
}

func TestGetListForPVZ_Success(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	expectedReceptions := []models.Reception{
		{ID: uuid.New(), DateTime: time.Now(), PvzID: pvzID, Status: "open"},
		{ID: uuid.New(), DateTime: time.Now(), PvzID: pvzID, Status: "closed"},
	}

	receptionRepo.On("GetAllForPVZ", ctx, pvzID).Return(expectedReceptions, nil)

	useCase := NewReceptionUseCase(receptionRepo, logger)

	receptions, err := useCase.GetListForPVZ(ctx, pvzID)

	assert.NoError(t, err)
	assert.Len(t, receptions, 2)
	assert.Equal(t, expectedReceptions, receptions)
	receptionRepo.AssertCalled(t, "GetAllForPVZ", ctx, pvzID)
}

func TestGetListForPVZ_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("GetAllForPVZ", ctx, pvzID).Return(nil, errors.New("failed to fetch reception list"))

	useCase := NewReceptionUseCase(receptionRepo, logger)

	receptions, err := useCase.GetListForPVZ(ctx, pvzID)

	assert.Error(t, err)
	assert.Nil(t, receptions)
	receptionRepo.AssertCalled(t, "GetAllForPVZ", ctx, pvzID)
}
