package usecase

import (
	"avito-pvz/internal/models"
	"avito-pvz/internal/usecase/reception"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

// ReceptionUseCase contains business logic for managing reception operations
type ReceptionUseCase struct {
	receptionRepo reception.Repository
	logger        *slog.Logger
}

// NewReceptionUseCase creates a new ReceptionUseCase
func NewReceptionUseCase(rRepo reception.Repository, logger *slog.Logger) *ReceptionUseCase {
	return &ReceptionUseCase{
		receptionRepo: rRepo,
		logger:        logger,
	}
}

// StartReception starts a new reception by adding it to the repository (if not open)
func (r *ReceptionUseCase) StartReception(ctx context.Context, reception models.Reception) error {
	r.logger.Info("Attempting to start reception", "receptionID", reception.ID)

	if err := r.receptionRepo.Create(ctx, reception); err != nil {
		r.logger.Error("Failed to start reception", "receptionID", reception.ID, "error", err)
		return err
	}

	r.logger.Info("Reception started successfully", "receptionID", reception.ID)
	return nil
}

// CloseReception closes the specified reception for the given PVZ (if open)
func (r *ReceptionUseCase) CloseReception(ctx context.Context, receptionID uuid.UUID, pvzId uuid.UUID) error {
	r.logger.Info("Attempting to close reception", "receptionID", receptionID)

	isOpen := r.receptionRepo.IsOpen(ctx, receptionID)
	if !isOpen {
		r.logger.Info("Reception already closed", "receptionID", receptionID)
		return errors.New("reception is already closed")
	}

	err := r.receptionRepo.CloseLast(ctx, pvzId)
	if err != nil {
		r.logger.Error("Failed to close reception", "receptionID", receptionID, "error", err)
		return err
	}

	r.logger.Info("Reception closed successfully", "receptionID", receptionID)
	return nil
}

// GetLastReception retrieves the most recent reception for the given PVZ ID
func (r *ReceptionUseCase) GetLastReception(ctx context.Context, pvzID uuid.UUID) (models.Reception, error) {
	r.logger.Info("Attempting to get last reception for PVZ", "pvzID", pvzID)

	MyReception, err := r.receptionRepo.GetLast(ctx, pvzID)
	if err != nil {
		r.logger.Error("Failed to get last reception", "pvzID", pvzID, "error", err)
		return models.Reception{}, err
	}

	r.logger.Info("Successfully retrieved last reception", "pvzID", pvzID, "receptionID", MyReception.ID)
	return MyReception, nil
}

// GetListForPVZ retrieves all receptions for the specified PVZ ID
func (r *ReceptionUseCase) GetListForPVZ(ctx context.Context, pvzID uuid.UUID) ([]models.Reception, error) {
	r.logger.Info("Attempting to get reception list for PVZ", "pvzID", pvzID)

	MyReceptions, err := r.receptionRepo.GetAllForPVZ(ctx, pvzID)
	if err != nil {
		r.logger.Error("Failed to get reception list", "pvzID", pvzID, "error", err)
		return nil, err
	}

	r.logger.Info("Successfully retrieved reception list", "pvzID", pvzID)
	return MyReceptions, nil
}
