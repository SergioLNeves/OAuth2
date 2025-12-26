package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
)

type Config struct {
	DBPath      string
	Environment string
	MaxConn     int
	MaxIdle     int
	MaxLifeTime time.Duration
}

type sqliteDB struct {
	db *gorm.DB
}

func NewDatabase(cfg *Config) (ports.Database, error) {
	if cfg == nil {
		return nil, fmt.Errorf("database config cannot be nil")
	}

	dbDir := filepath.Dir(cfg.DBPath)
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	logLevel := logger.Info
	if cfg.Environment == "production" {
		logLevel = logger.Warn
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	}

	db, err := gorm.Open(sqlite.Open(cfg.DBPath), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Successfully connected to SQLite database at %s", cfg.DBPath)

	return &sqliteDB{db: db}, nil
}

func (s *sqliteDB) GetDB() *gorm.DB {
	return s.db
}

func (s *sqliteDB) Ping() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func (s *sqliteDB) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}

	log.Println("Database connection closed")
	return nil
}
