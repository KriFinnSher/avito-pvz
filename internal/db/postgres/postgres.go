package postgres

import (
	"avito-pvz/internal/config"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitDB initialize database with AppConfig's parameters for a defined driver
func InitDB() (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.User,
		config.AppConfig.DB.Pass,
		config.AppConfig.DB.Name,
	)
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MakeMigrations use all *.up.sql files if up is true, and *.down.sql otherwise
func MakeMigrations(up bool) error {
	dbLine := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable",
		config.AppConfig.DB.User,
		config.AppConfig.DB.Pass,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.Name,
	)
	m, err := migrate.New("file://migrations", dbLine)
	if err != nil {
		return err
	}

	if up {
		err = m.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return nil
			}
			return err
		}
	} else {
		err = m.Down()
		if err != nil {
			return err
		}
	}
	return nil
}
