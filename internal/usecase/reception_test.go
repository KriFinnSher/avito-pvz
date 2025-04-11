package usecase

import (
	mockReception "avito-pvz/internal/mocks/reception"
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

func TestStartReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример нового Reception
	newReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "open",
	}

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(false) // прием еще не открыт
	receptionRepo.On("Create", ctx, newReception).Return(nil)  // успешное создание приема

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	err := useCase.StartReception(ctx, newReception)

	// Проверки
	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "Create", ctx, newReception)
}

func TestStartReception_Error_AlreadyOpen(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример нового Reception
	newReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "open",
	}

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(true) // прием уже открыт

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	err := useCase.StartReception(ctx, newReception)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "reception is already open", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	receptionRepo.AssertNotCalled(t, "Create", ctx, newReception)
}

func TestCloseReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(true) // прием открыт
	receptionRepo.On("CloseLast", ctx, pvzID).Return(nil)     // успешное закрытие приема

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	err := useCase.CloseReception(ctx, receptionID, pvzID)

	// Проверки
	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	receptionRepo.AssertCalled(t, "CloseLast", ctx, pvzID)
}

func TestCloseReception_Error_AlreadyClosed(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(false) // прием уже закрыт

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	err := useCase.CloseReception(ctx, receptionID, pvzID)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "reception is already closed", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	receptionRepo.AssertNotCalled(t, "CloseLast", ctx, pvzID)
}

func TestGetLastReception_Success(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()
	receptionID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример последнего Reception
	expectedReception := models.Reception{
		ID:       receptionID,
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   "closed",
	}

	// Настройка поведения моков
	receptionRepo.On("GetLast", ctx, pvzID).Return(expectedReception, nil) // успешное получение последнего приема

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	reception, err := useCase.GetLastReception(ctx, pvzID)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, expectedReception, reception)
	receptionRepo.AssertCalled(t, "GetLast", ctx, pvzID)
}

func TestGetLastReception_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("GetLast", ctx, pvzID).Return(models.Reception{}, errors.New("failed to fetch last reception")) // ошибка при получении последнего приема

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	reception, err := useCase.GetLastReception(ctx, pvzID)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "failed to fetch last reception", err.Error())
	assert.Equal(t, models.Reception{}, reception)
	receptionRepo.AssertCalled(t, "GetLast", ctx, pvzID)
}

func TestGetListForPVZ_Success(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Пример списка Reception
	expectedReceptions := []models.Reception{
		{ID: uuid.New(), DateTime: time.Now(), PvzID: pvzID, Status: "open"},
		{ID: uuid.New(), DateTime: time.Now(), PvzID: pvzID, Status: "closed"},
	}

	// Настройка поведения моков
	receptionRepo.On("GetAllForPVZ", ctx, pvzID).Return(expectedReceptions, nil) // успешное получение списка приемов

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	receptions, err := useCase.GetListForPVZ(ctx, pvzID)

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, receptions, 2)                    // проверяем, что вернулось два приема
	assert.Equal(t, expectedReceptions, receptions) // проверяем, что возвращенный список соответствует ожидаемому
	receptionRepo.AssertCalled(t, "GetAllForPVZ", ctx, pvzID)
}

func TestGetListForPVZ_Error(t *testing.T) {
	ctx := context.Background()
	pvzID := uuid.New()

	// Моки
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("GetAllForPVZ", ctx, pvzID).Return(nil, errors.New("failed to fetch reception list")) // ошибка при получении списка приемов

	// Создание UseCase
	useCase := NewReceptionUseCase(receptionRepo, logger)

	// Вызов метода
	receptions, err := useCase.GetListForPVZ(ctx, pvzID)

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, receptions) // если ошибка, возвращаем nil
	receptionRepo.AssertCalled(t, "GetAllForPVZ", ctx, pvzID)
}
