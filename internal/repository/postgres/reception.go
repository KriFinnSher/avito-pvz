package postgres

import (
	"avito-pvz/internal/models"
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type ReceptionRepo struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewReceptionRepo(db *sqlx.DB, logger *slog.Logger) *ReceptionRepo {
	return &ReceptionRepo{
		db:     db,
		logger: logger,
	}
}

func (r *ReceptionRepo) Create(ctx context.Context, reception models.Reception) error {
	query, args, err := sq.
		Insert("receptions").
		Columns("id", "date_time", "pvz_id", "status").
		Values(reception.ID, reception.DateTime, reception.PvzID, reception.Status).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build INSERT query for reception", "error", err, "reception_id", reception.ID, "pvz_id", reception.PvzID)
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Error("failed to insert reception", "error", err, "reception_id", reception.ID, "pvz_id", reception.PvzID)
		return err
	}

	r.logger.Info("reception created successfully", "reception_id", reception.ID, "pvz_id", reception.PvzID)
	return nil
}

func (r *ReceptionRepo) CloseLast(ctx context.Context, pvzID uuid.UUID) error {
	query, args, err := sq.
		Select("id").
		From("receptions").
		Where(sq.Eq{"pvz_id": pvzID}).
		OrderBy("date_time DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for last reception", "error", err, "pvz_id", pvzID)
		return err
	}

	var lastReceptionID uuid.UUID
	err = r.db.GetContext(ctx, &lastReceptionID, query, args...)
	if err != nil {
		r.logger.Error("failed to fetch last reception", "error", err, "pvz_id", pvzID)
		return err
	}

	updateQuery, updateArgs, updateErr := sq.
		Update("receptions").
		Set("status", "close").
		Where(sq.Eq{"id": lastReceptionID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if updateErr != nil {
		r.logger.Error("failed to build UPDATE query for closing reception", "error", updateErr, "reception_id", lastReceptionID)
		return updateErr
	}

	_, err = r.db.ExecContext(ctx, updateQuery, updateArgs...)
	if err != nil {
		r.logger.Error("failed to close reception", "error", err, "reception_id", lastReceptionID)
		return err
	}

	r.logger.Info("last reception closed successfully", "reception_id", lastReceptionID, "pvz_id", pvzID)
	return nil
}

func (r *ReceptionRepo) IsOpen(ctx context.Context, receptionID uuid.UUID) bool {
	query, args, err := sq.Select("status").
		From("receptions").
		Where(sq.Eq{"id": receptionID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for reception status check", "error", err, "receptionID", receptionID)
		return false
	}

	var status string
	err = r.db.GetContext(ctx, &status, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Info("reception not found", "receptionID", receptionID)
			return false
		}
		r.logger.Error("failed to check reception status", "error", err, "receptionID", receptionID)
		return false
	}

	if status == "in_progress" {
		r.logger.Info("reception is open", "receptionID", receptionID)
		return true
	}

	r.logger.Info("reception is closed", "receptionID", receptionID)
	return false
}

func (r *ReceptionRepo) GetLast(ctx context.Context, pvzID uuid.UUID) (models.Reception, error) {
	query, args, err := sq.Select("id", "date_time", "pvz_id", "status").
		From("receptions").
		Where(sq.Eq{"pvz_id": pvzID}).
		OrderBy("date_time DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for reception fetching", "error", err, "pvzID", pvzID)
		return models.Reception{}, err
	}

	var reception models.Reception
	err = r.db.GetContext(ctx, &reception, query, args...)
	if err != nil {
		r.logger.Error("failed to fetch last reception", "error", err, "pvz_id", pvzID)
		return models.Reception{}, err
	}

	return reception, nil
}

func (r *ReceptionRepo) GetAllForPVZ(ctx context.Context, pvzID uuid.UUID) ([]models.Reception, error) {
	query, args, err := sq.
		Select("id", "date_time", "pvz_id", "status").
		From("receptions").
		Where(sq.Eq{"pvz_id": pvzID}).
		OrderBy("date_time DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for reception list", "error", err, "pvzID", pvzID)
		return nil, err
	}

	var receptionList []models.Reception
	err = r.db.SelectContext(ctx, &receptionList, query, args...)
	if err != nil {
		r.logger.Error("failed to fetch reception list", "error", err, "pvzID", pvzID)
		return nil, err
	}

	r.logger.Info("fetched reception list successfully", "count", len(receptionList))
	return receptionList, nil
}
