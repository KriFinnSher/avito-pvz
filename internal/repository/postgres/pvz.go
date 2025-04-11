package postgres

import (
	"avito-pvz/internal/models"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

// PVZRepo manages database operations related to pickup points (pvz)
type PVZRepo struct {
	db     *sqlx.DB
	logger *slog.Logger
}

// NewPVZRepo creates a new instance of PVZRepo
func NewPVZRepo(db *sqlx.DB, logger *slog.Logger) *PVZRepo {
	return &PVZRepo{
		db:     db,
		logger: logger,
	}
}

// Create inserts a new PVZ record into the database
func (p *PVZRepo) Create(ctx context.Context, pvz models.PVZ) error {
	query, args, err := sq.
		Insert("pvz").
		Columns("id", "city").
		Values(pvz.ID, pvz.City).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build INSERT query for PVZ", "error", err, "pvz_id", pvz.ID, "city", pvz.City)
		return err
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Error("failed to insert PVZ", "error", err, "pvz_id", pvz.ID, "city", pvz.City)
		return err
	}

	p.logger.Info("pvz created successfully", "pvz_id", pvz.ID, "city", pvz.City)
	return nil
}

// GetAll retrieves all PVZ records from the database
func (p *PVZRepo) GetAll(ctx context.Context) ([]models.PVZ, error) {
	query, args, err := sq.
		Select("id", "registration_date", "city").
		From("pvz").
		OrderBy("registration_date DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build SELECT query for PVZ list", "error", err)
		return nil, err
	}

	var pvzList []models.PVZ
	err = p.db.SelectContext(ctx, &pvzList, query, args...)
	if err != nil {
		p.logger.Error("failed to fetch PVZ list", "error", err)
		return nil, err
	}

	p.logger.Info("fetched PVZ list successfully", "count", len(pvzList))
	return pvzList, nil
}
