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

	// Моки
	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример нового PVZ
	newPVZ := models.PVZ{
		ID:               pvzID,
		RegistrationDate: time.Now(),
		City:             "Test City",
	}

	// Настройка поведения моков
	pvzRepo.On("Create", ctx, newPVZ).Return(nil) // успешное создание PVZ

	// Создание UseCase
	useCase := NewPVZUseCase(pvzRepo, logger)

	// Вызов метода
	err := useCase.CreatePVZ(ctx, newPVZ)

	// Проверки
	assert.NoError(t, err)
	pvzRepo.AssertCalled(t, "Create", ctx, newPVZ)
}

func TestCreatePVZ_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	// Моки
	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример нового PVZ
	newPVZ := models.PVZ{
		ID:               pvzID,
		RegistrationDate: time.Now(),
		City:             "Test City",
	}

	// Настройка поведения моков
	pvzRepo.On("Create", ctx, newPVZ).Return(errors.New("failed to create PVZ")) // ошибка при создании PVZ

	// Создание UseCase
	useCase := NewPVZUseCase(pvzRepo, logger)

	// Вызов метода
	err := useCase.CreatePVZ(ctx, newPVZ)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "failed to create PVZ", err.Error())
	pvzRepo.AssertCalled(t, "Create", ctx, newPVZ)
}

func TestGetAllPVZs_Success(t *testing.T) {
	ctx := context.Background()

	// Моки
	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример списка PVZ
	expectedPVZs := []models.PVZ{
		{ID: uuid.New(), RegistrationDate: time.Now(), City: "City 1"},
		{ID: uuid.New(), RegistrationDate: time.Now(), City: "City 2"},
	}

	// Настройка поведения моков
	pvzRepo.On("GetAll", ctx).Return(expectedPVZs, nil) // успешное получение списка PVZ

	// Создание UseCase
	useCase := NewPVZUseCase(pvzRepo, logger)

	// Вызов метода
	pvzs, err := useCase.GetAllPVZs(ctx)

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, pvzs, 2)              // проверяем, что вернулось два PVZ
	assert.Equal(t, expectedPVZs, pvzs) // проверяем, что возвращенный список соответствует ожидаемому
	pvzRepo.AssertCalled(t, "GetAll", ctx)
}

func TestGetAllPVZs_Error(t *testing.T) {
	ctx := context.Background()

	// Моки
	pvzRepo := mockPVZ.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	pvzRepo.On("GetAll", ctx).Return(nil, errors.New("failed to fetch PVZs")) // ошибка при получении списка PVZ

	// Создание UseCase
	useCase := NewPVZUseCase(pvzRepo, logger)

	// Вызов метода
	pvzs, err := useCase.GetAllPVZs(ctx)

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, pvzs) // если ошибка, возвращаем nil
	pvzRepo.AssertCalled(t, "GetAll", ctx)
}
