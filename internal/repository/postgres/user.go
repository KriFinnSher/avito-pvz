package postgres

import (
	"avito-pvz/internal/models"
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type UserRepo struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewUserRepo(db *sqlx.DB, logger *slog.Logger) *UserRepo {
	return &UserRepo{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, user models.User) error {
	query, args, err := sq.
		Insert("users").
		Columns("id", "email", "role", "hash").
		Values(user.ID, user.Email, user.Role, user.Hash).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build INSERT query for user", "error", err, "user_id", user.ID, "email", user.Email)
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Error("failed to insert user", "error", err, "user_id", user.ID, "email", user.Email)
		return err
	}

	r.logger.Info("user created successfully", "user_id", user.ID, "email", user.Email)
	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	query, args, err := sq.
		Select("id", "email", "role", "hash").
		From("users").
		Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for user by email", "error", err, "email", email)
		return models.User{}, err
	}

	var user models.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Info("user not found", "email", email)
			return models.User{}, nil
		}
		r.logger.Error("failed to fetch user by email", "error", err, "email", email)
		return models.User{}, err
	}

	r.logger.Info("user fetched successfully", "user_id", user.ID, "email", user.Email)
	return user, nil
}

func (r *UserRepo) Exists(ctx context.Context, email string) (bool, error) {
	query, args, err := sq.
		Select("1").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		r.logger.Error("failed to build SELECT query for user existence by email", "error", err, "email", email)
		return false, err
	}

	var exists bool
	err = r.db.GetContext(ctx, &exists, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Info("user not found", "email", email)
			return false, nil
		}
		r.logger.Error("failed to check user existence by email", "error", err, "email", email)
		return false, err
	}

	r.logger.Info("user exists", "email", email)
	return true, nil
}
