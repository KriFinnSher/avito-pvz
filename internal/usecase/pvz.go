package usecase

import (
	"avito-pvz/internal/models"
	"avito-pvz/internal/usecase/pvz"
	"context"
	"log/slog"
)

type PVZUseCase struct {
	pvzRepo pvz.Repository
	logger  *slog.Logger
}

func NewPVZUseCase(pRepo pvz.Repository, logger *slog.Logger) *PVZUseCase {
	return &PVZUseCase{
		pvzRepo: pRepo,
		logger:  logger,
	}
}

func (p *PVZUseCase) CreatePVZ(ctx context.Context, pvz models.PVZ) error {
	p.logger.Info("Attempting to create PVZ", "pvz_id", pvz.ID)

	if err := p.pvzRepo.Create(ctx, pvz); err != nil {
		p.logger.Error("Failed to create PVZ", "pvz_id", pvz.ID, "error", err)
		return err
	}

	p.logger.Info("Successfully created PVZ", "pvz_id", pvz.ID)
	return nil
}

func (p *PVZUseCase) GetAllPVZs(ctx context.Context) ([]models.PVZ, error) {
	p.logger.Info("Fetching all PVZs")

	pvzs, err := p.pvzRepo.GetAll(ctx)
	if err != nil {
		p.logger.Error("Failed to fetch PVZs", "error", err)
		return nil, err
	}

	p.logger.Info("Successfully fetched all PVZs", "count", len(pvzs))
	return pvzs, nil
}
