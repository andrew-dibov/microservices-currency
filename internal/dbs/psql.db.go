package dbs

import (
	"context"
	"database/sql"
	"fmt"
	"microservices-currency/internal/configs"
)

func NewPsqlDB(c *configs.AppConfig) (*PsqlDB, error) {
	db, err := sql.Open("postgres", c.Infra.Psql)
	if err != nil {
		return nil, fmt.Errorf("failed to open : %w", err)
	}

	db.SetMaxOpenConns(c.Infra.PsqlDbMaxOpen)
	db.SetMaxIdleConns(c.Infra.PsqlDbMaxIdle)
	db.SetConnMaxLifetime(c.Infra.PsqlDbMaxLifetime)

	ctx, can := context.WithTimeout(context.Background(), c.Timeouts.Psql)
	defer can()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping : %w", err)
	}

	return &PsqlDB{db}, nil
}

func (db *PsqlDB) Close() error {
	return db.DB.Close()
}
