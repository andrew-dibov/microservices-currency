package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

func NewPostgresRepository(postgresDatabase *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		postgresDatabase: postgresDatabase,
	}
}

func (repository *PostgresRepository) GetRate(ctx context.Context, fromCurrency string, toCurrency string) (float64, error) {
	query := `
	SELECT rate FROM rates
	WHERE currency_code = $1
	AND base_currency = 'USD'
	AND updated_at > NOW() - INTERVAL '1 day'
	`

	/* --- --- --- */

	if fromCurrency == toCurrency {
		return 1.00, nil
	}

	/* --- --- --- */

	var rate float64

	if fromCurrency == "USD" {
		err := repository.postgresDatabase.QueryRowContext(ctx, query, toCurrency).Scan(&rate)
		if err != nil {
			return 0, err
		}
		return rate, nil
	}

	if toCurrency == "USD" {
		err := repository.postgresDatabase.QueryRowContext(ctx, query, fromCurrency).Scan(&rate)
		if err != nil {
			return 0, err
		}
		return 1 / rate, nil
	}

	/* --- --- --- */

	var fromRate, toRate float64

	err := repository.postgresDatabase.QueryRowContext(ctx, query, fromCurrency).Scan(&fromRate)
	if err != nil {
		return 0, err
	}

	err = repository.postgresDatabase.QueryRowContext(ctx, query, toCurrency).Scan(&toRate)
	if err != nil {
		return 0, err
	}

	return toRate / fromRate, nil
}

func (repository *PostgresRepository) GetRates(ctx context.Context, baseCurrency string) (map[string]float64, error) {
	query := `
	SELECT currency_code, rate FROM rates
	WHERE base_currency = $1
	AND updated_at > NOW() - INTERVAL '1 day'
	`

	/* --- --- --- */

	rows, err := repository.postgresDatabase.QueryContext(ctx, query, baseCurrency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/* --- --- --- */

	rates := make(map[string]float64)
	for rows.Next() {
		var code string
		var rate float64

		if err := rows.Scan(&code, &rate); err != nil {
			return nil, err
		}
		rates[code] = rate
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(rates) == 0 {
		return nil, sql.ErrNoRows
	}

	/* --- --- --- */

	return rates, nil
}

func (repository *PostgresRepository) UpdateRates(ctx context.Context, baseCurrency string, rates map[string]float64) error {
	if len(rates) == 0 {
		return fmt.Errorf("rates is empty")
	}

	/* --- --- --- */

	tx, err := repository.postgresDatabase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	/* --- --- --- */

	if _, err := tx.ExecContext(ctx, "DELETE FROM rates WHERE base_currency = $1", baseCurrency); err != nil {
		return err
	}

	query := `
	INSERT INTO rates (base_currency, currency_code, rate, updated_at)
	VALUES ($1, $2, $3, NOW())
	`

	for code, rate := range rates {
		if _, err := tx.ExecContext(ctx, query, baseCurrency, code, rate); err != nil {
			return err
		}
	}

	return tx.Commit()
}
