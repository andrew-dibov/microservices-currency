package configs

import "time"

type AppConfig struct {
	App App

	PostgresDatabase PostgresDatabase
	ExchangeClient   ExchangeClient

	Updater Updater
}

/* --- --- --- */

type App struct {
	Name string

	Prod bool
	Port string

	KeepaliveTime    time.Duration
	KeepaliveTimeout time.Duration

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

/* --- --- --- */

type Updater struct {
	BaseCurrency string

	UpdateInterval time.Duration

	FetchTimeout time.Duration
	StoreTimeout time.Duration
}
