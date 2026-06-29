package database

import (
	"log/slog"
	"time"

	"mintrix-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect opens a PostgreSQL connection via GORM, configures the connection pool,
// and runs auto-migration for development.
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// AutoMigrate for development only — production must use golang-migrate.
	if err := db.AutoMigrate(&models.User{}, &models.RefreshToken{}); err != nil {
		return nil, err
	}

	slog.Info("Database connected and migrated")
	return db, nil
}
