package usecase

import (
	"avito-pvz/internal/models"
	mockProduct "avito-pvz/internal/repository/mocks/product"
	mockReception "avito-pvz/internal/repository/mocks/reception"
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"log/slog"
)

func TestRemoveFromReception_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(true)
	productRepo.On("DeleteLast", ctx, receptionID).Return(nil)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	err := useCase.RemoveFromReception(ctx, receptionID)

	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertCalled(t, "DeleteLast", ctx, receptionID)
}

func TestRemoveFromReception_ReceptionClosed_Error(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(false)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	err := useCase.RemoveFromReception(ctx, receptionID)

	assert.Error(t, err)
	assert.Equal(t, "reception is closed, unable to interact", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertNotCalled(t, "DeleteLast", ctx, receptionID)
}

func TestAddNew_Success(t *testing.T) {
	ctx := context.Background()
	productID := uuid.New()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(true)
	productRepo.On("AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID}).Return(nil)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	err := useCase.AddNew(ctx, models.Product{ID: productID, ReceptionId: receptionID})

	assert.NoError(t, err)
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertCalled(t, "AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID})
}

func TestAddNew_ReceptionClosed_Error(t *testing.T) {
	ctx := context.Background()
	productID := uuid.New()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	receptionRepo.On("IsOpen", ctx, receptionID).Return(false)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	err := useCase.AddNew(ctx, models.Product{ID: productID, ReceptionId: receptionID})

	assert.Error(t, err)
	assert.Equal(t, "reception is closed, unable to interact", err.Error())
	receptionRepo.AssertCalled(t, "IsOpen", ctx, receptionID)
	productRepo.AssertNotCalled(t, "AddOne", ctx, models.Product{ID: productID, ReceptionId: receptionID})
}

func TestGetReceptionList_Success(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	productRepo.On("GetForReception", ctx, receptionID).Return([]models.Product{{ID: uuid.New()}}, nil)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	products, err := useCase.GetReceptionList(ctx, receptionID)

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	productRepo.AssertCalled(t, "GetForReception", ctx, receptionID)
}

func TestGetReceptionList_Error(t *testing.T) {
	ctx := context.Background()
	receptionID := uuid.New()

	productRepo := mockProduct.NewRepository(t)
	receptionRepo := mockReception.NewRepository(t)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	productRepo.On("GetForReception", ctx, receptionID).Return(nil, assert.AnError)

	useCase := NewProductUseCase(productRepo, receptionRepo, logger)

	products, err := useCase.GetReceptionList(ctx, receptionID)

	assert.Error(t, err)
	assert.Nil(t, products)
	productRepo.AssertCalled(t, "GetForReception", ctx, receptionID)
}
