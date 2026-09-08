package configs

import "time"

type AppConfig struct {
	App App

	PostgresDatabase PostgresDatabase
	ExchangeClient   ExchangeClient
}

/* --- --- --- */

type App struct {
	Name string

	Prod bool
	Port string

	ReadTimeout     time.Duration
	IdleTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

/* --- --- --- */

type PostgresDatabase struct {
	Address string
}

/* --- --- --- */

type ExchangeClient struct {
	Address string
	Token   string

	Timeout time.Duration
}
