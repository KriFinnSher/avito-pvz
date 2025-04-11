package usecase

import (
	"avito-pvz/internal/models"
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockProduct "avito-pvz/internal/mocks/product"
	mockReception "avito-pvz/internal/mocks/reception"

	"log/slog"
)

func TestRemoveFromReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(true)  // приемка открыта
	productRepo.On("DeleteLast", ctx, receptionID).Return(nil) // успешное удаление продукта

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	err := useCase.RemoveFromReception(ctx, receptionID)

	// Проверки
	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertCalled(t, "DeleteLast", ctx, receptionID)
}

func TestRemoveFromReception_ReceptionClosed_Error(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(false) // приемка закрыта

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	err := useCase.RemoveFromReception(ctx, receptionID)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "reception is closed, unable to interact", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertNotCalled(t, "DeleteLast", ctx, receptionID)
}

func TestAddNew_Success(t *testing.T) {
	ctx := context.Background()
	productID := uuid.New()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(true)                                          // приемка открыта
	productRepo.On("AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID}).Return(nil) // успешное добавление продукта

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	err := useCase.AddNew(ctx, models.Product{ID: productID, ReceptionId: receptionID})

	// Проверки
	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertCalled(t, "AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID})
}

func TestAddNew_ReceptionClosed_Error(t *testing.T) {
	ctx := context.Background()
	productID := uuid.New()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	receptionRepo.On("IsOpen", ctx, receptionID).Return(false) // приемка закрыта

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	err := useCase.AddNew(ctx, models.Product{ID: productID, ReceptionId: receptionID})

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "reception is closed, unable to interact", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertNotCalled(t, "AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID})
}

func TestGetReceptionList_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	productRepo.On("GetForReception", ctx, receptionID).Return([]models.Product{{ID: uuid.New()}}, nil) // успешное получение списка продуктов

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	products, err := useCase.GetReceptionList(ctx, receptionID)

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, products, 1) // ожидаем, что вернется 1 продукт
	productRepo.AssertCalled(t, "GetForReception", ctx, receptionID)
}

func TestGetReceptionList_Error(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	// Моки
	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Настройка поведения моков
	productRepo.On("GetForReception", ctx, receptionID).Return(nil, assert.AnError) // ошибка при получении продуктов

	// Создание UseCase
	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	// Вызов метода
	products, err := useCase.GetReceptionList(ctx, receptionID)

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, products)
	productRepo.AssertCalled(t, "GetForReception", ctx, receptionID)
}
