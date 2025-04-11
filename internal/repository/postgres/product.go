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

// ProductRepo handles database operations related to products table
type ProductRepo struct {
	db     *sqlx.DB
	logger *slog.Logger
}

// NewProductRepo creates a new instance of ProductRepo
func NewProductRepo(db *sqlx.DB, logger *slog.Logger) *ProductRepo {
	return &ProductRepo{
		db:     db,
		logger: logger,
	}
}

// DeleteLast removes the last added product for the reception with specific id
func (p *ProductRepo) DeleteLast(ctx context.Context, receptionID uuid.UUID) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		p.logger.Error("failed to begin transaction", "error", err, "reception_id", receptionID)
		return err
	}
	defer tx.Rollback()

	query, args, err := sq.
		Select("id").
		From("products").
		Where(sq.Eq{"reception_id": receptionID}).
		OrderBy("date_time DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build SELECT query", "error", err, "reception_id", receptionID)
		return err
	}

	var productID uuid.UUID
	err = tx.GetContext(ctx, &productID, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.logger.Warn("no product to delete", "reception_id", receptionID)
			return err
		}
		p.logger.Error("failed to execute SELECT query", "error", err, "reception_id", receptionID)
		return err
	}

	delQuery, delArgs, err := sq.
		Delete("products").
		Where(sq.Eq{"id": productID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build DELETE query", "error", err, "product_id", productID)
		return err
	}

	_, err = tx.ExecContext(ctx, delQuery, delArgs...)
	if err != nil {
		p.logger.Error("failed to execute DELETE", "error", err, "product_id", productID)
		return err
	}

	if err := tx.Commit(); err != nil {
		p.logger.Error("failed to commit transaction", "error", err)
		return err
	}

	p.logger.Info("product deleted successfully", "product_id", productID, "reception_id", receptionID)
	return nil
}

// AddOne inserts a new product into the repository
func (p *ProductRepo) AddOne(ctx context.Context, product models.Product) error {
	query, args, err := sq.
		Insert("products").
		Columns("id", "type", "reception_id").
		Values(product.ID, product.Type, product.ReceptionId).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build INSERT query", "error", err, "product_id", product.ID, "reception_id", product.ReceptionId)
		return err
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	if err != nil {
		p.logger.Error("failed to insert product", "error", err, "product_id", product.ID, "reception_id", product.ReceptionId)
		return err
	}

	p.logger.Info("product inserted successfully", "product_id", product.ID, "reception_id", product.ReceptionId, "type", product.Type)
	return nil
}

// GetForReception returns products linked to the reception with specific id
func (p *ProductRepo) GetForReception(ctx context.Context, receptionID uuid.UUID) ([]models.Product, error) {
	query, args, err := sq.
		Select("id", "date_time", "type", "reception_id").
		From("products").
		Where(sq.Eq{"reception_id": receptionID}).
		OrderBy("date_time ASC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		p.logger.Error("failed to build SELECT query", "error", err, "reception_id", receptionID)
		return nil, err
	}

	var products []models.Product
	err = p.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		p.logger.Error("failed to fetch products for reception", "error", err, "reception_id", receptionID)
		return nil, err
	}

	p.logger.Info("fetched products for reception", "count", len(products), "reception_id", receptionID)
	return products, nil
}
