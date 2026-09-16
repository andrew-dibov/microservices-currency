CREATE TABLE IF NOT EXISTS rates (
    id            SERIAL PRIMARY KEY,
    base_currency VARCHAR(3)      NOT NULL,
    currency_code VARCHAR(3)      NOT NULL,
    rate          NUMERIC(20, 10) NOT NULL,
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_base_code UNIQUE (base_currency, currency_code)
);

CREATE INDEX IF NOT EXISTS idx_rates_base_updated
    ON rates (base_currency, updated_at DESC);
