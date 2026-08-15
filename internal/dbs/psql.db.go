package dbs

import (
	"context"
	"database/sql"
	"fmt"
	"microservices-currency/internal/configs"
)

func NewPsqlDb(cfg *configs.AppConfig) (*PsqlDb, error) {
	db, err := sql.Open("postgres", cfg.Infra.Psql)
	if err != nil {
		return nil, fmt.Errorf("failed to open %w", err)
	}

	db.SetMaxOpenConns(cfg.Infra.PsqlDbMaxOpen)
	db.SetMaxIdleConns(cfg.Infra.PsqlDbMaxIdle)
	db.SetConnMaxLifetime(cfg.Infra.PsqlDbMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeouts.Psql)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping %w", err)
	}

	return &PsqlDb{db}, nil
}

func (db *PsqlDb) Close() error {
	return db.DB.Close()
}
