package repos

import (
	"context"
	"database/sql"
	"fmt"
)

func NewPsqlRepo(psql *sql.DB) *PsqlRepo {
	return &PsqlRepo{
		psql: psql,
	}
}

func (r *PsqlRepo) Rate(ctx context.Context, fromCurrency string, toCurrency string) (float64, error) {
	if fromCurrency == toCurrency {
		return 1.00, nil
	}

	var rate float64

	if fromCurrency == "USD" {
		err := r.psql.QueryRowContext(ctx,
			`SELECT rate FROM rates WHERE base_currency = 'USD' AND currency_code = $1`,
			toCurrency,
		).Scan(&rate)

		if err != nil {
			return 0, err
		}
		return rate, nil
	}

	if toCurrency == "USD" {
		err := r.psql.QueryRowContext(ctx,
			`SELECT rate FROM rates WHERE base_currency = 'USD' AND currency_code = $1`,
			fromCurrency,
		).Scan(&rate)

		if err != nil {
			return 0, err
		}
		return 1 / rate, nil
	}

	var fromRate, toRate float64

	err := r.psql.QueryRowContext(ctx,
		`SELECT rate FROM rates WHERE base_currency = 'USD' AND currency_code = $1`,
		fromCurrency,
	).Scan(&fromRate)

	if err != nil {
		return 0, err
	}

	err = r.psql.QueryRowContext(ctx,
		`SELECT rate FROM rates WHERE base_currency = 'USD' AND currency_code = $1`,
		toCurrency,
	).Scan(&toRate)

	if err != nil {
		return 0, err
	}

	return toRate / fromRate, nil
}

func (r *PsqlRepo) Rates(ctx context.Context, baseCurrency string) (map[string]float64, error) {
	if baseCurrency == "" {
		baseCurrency = "USD"
	}

	rows, err := r.psql.QueryContext(ctx,
		`SELECT currency_code, rate FROM rates WHERE base_currency = 'USD'`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ratesUSD := make(map[string]float64)

	for rows.Next() {
		var code string
		var rate float64

		if err := rows.Scan(&code, &rate); err != nil {
			return nil, err
		}
		ratesUSD[code] = rate
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(ratesUSD) == 0 {
		return nil, fmt.Errorf("no rates")
	}

	ratesUSD["USD"] = 1.0

	if baseCurrency == "USD" {
		return ratesUSD, nil
	}

	baseToUSD, ok := ratesUSD[baseCurrency]
	if !ok {
		return nil, fmt.Errorf("baseCurrency is absent")
	}

	res := make(map[string]float64, len(ratesUSD))
	for code, rateToUSD := range ratesUSD {
		res[code] = rateToUSD / baseToUSD
	}

	return res, nil
}

func (r *PsqlRepo) Update(ctx context.Context, baseCurrency string, rates map[string]float64) error {
	if baseCurrency == "" {
		return fmt.Errorf("baseCurrency is absent")
	}

	if len(rates) == 0 {
		return fmt.Errorf("rates is empty")
	}

	if len(rates) > 1000 {
		return fmt.Errorf("too many rates")
	}

	tx, err := r.psql.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
	INSERT INTO rates (base_currency, currency_code, rate, updated_at)
	VALUES ($1, $2, $3, NOW())
	ON CONFLICT (base_currency, currency_code)
	DO UPDATE SET rate = $3, updated_at = NOW()
	`

	for code, rate := range rates {
		if _, err := tx.ExecContext(ctx, q, baseCurrency, code, rate); err != nil {
			return err
		}
	}

	return tx.Commit()
}
