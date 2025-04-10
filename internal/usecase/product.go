package usecase

import (
	"avito-pvz/internal/models"
	"avito-pvz/internal/usecase/product"
	"avito-pvz/internal/usecase/reception"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type ProductUseCase struct {
	productRepo   product.Repository
	receptionRepo reception.Repository
	logger        *slog.Logger
}

func NewProductUseCase(pRepo product.Repository, rRepo reception.Repository, logger *slog.Logger) *ProductUseCase {
	return &ProductUseCase{
		productRepo:   pRepo,
		receptionRepo: rRepo,
		logger:        logger,
	}
}

func (p *ProductUseCase) RemoveFromReception(ctx context.Context, receptionID uuid.UUID) error {
	openStatus := p.receptionRepo.IsOpen(ctx, receptionID)
	if !openStatus {
		p.logger.Info("Attempted to remove product from closed reception", "receptionID", receptionID)
		return errors.New("reception is closed, unable to interact")
	}

	p.logger.Info("Proceeding to remove product from reception", "receptionID", receptionID)

	if err := p.productRepo.DeleteLast(ctx, receptionID); err != nil {
		p.logger.Error("Failed to delete last product from reception", "error", err, "receptionID", receptionID)
		return err
	}
	p.logger.Info("Successfully removed last product from reception", "receptionID", receptionID)
	return nil
}

func (p *ProductUseCase) AddNew(ctx context.Context, product models.Product) error {
	openStatus := p.receptionRepo.IsOpen(ctx, product.ReceptionId)
	if !openStatus {
		p.logger.Info("Attempted to add product to closed reception", "receptionID", product.ReceptionId)
		return errors.New("reception is closed, unable to interact")
	}

	p.logger.Info("Proceeding to add product to reception", "receptionID", product.ReceptionId, "productID", product.ID)

	if err := p.productRepo.AddOne(ctx, product); err != nil {
		p.logger.Error("Failed to add product to reception", "error", err, "receptionID", product.ReceptionId, "productID", product.ID)
		return err
	}

	p.logger.Info("Successfully added product to reception", "receptionID", product.ReceptionId, "productID", product.ID)
	return nil
}

func (p *ProductUseCase) GetReceptionList(ctx context.Context, receptionID uuid.UUID) ([]models.Product, error) {
	openStatus := p.receptionRepo.IsOpen(ctx, receptionID)
	if !openStatus {
		p.logger.Info("Attempted to get all products from closed reception", "receptionID", receptionID)
		return nil, errors.New("reception is closed, unable to interact")
	}

	p.logger.Info("Fetching products from reception", "receptionID", receptionID)

	products, err := p.productRepo.GetForReception(ctx, receptionID)
	if err != nil {
		p.logger.Error("Failed to fetch products for reception", "error", err, "receptionID", receptionID)
		return nil, err
	}

	p.logger.Info("Successfully fetched products from reception", "receptionID", receptionID, "productCount", len(products))
	return products, nil
}
