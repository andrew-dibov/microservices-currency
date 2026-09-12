package configs

import "time"

type AppConfig struct {
	App App

	PostgresDatabase PostgresDatabase
	ExchangeClient   ExchangeClient
	UpdaterJob       UpdaterJob
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
	DSN string
}

/* --- --- --- */

type ExchangeClient struct {
	Token   string
	Address string
	Timeout time.Duration
}

/* --- --- --- */

type UpdaterJob struct {
	UpdateInterval time.Duration
	FetchTimeout   time.Duration
	StoreTimeout   time.Duration
}
