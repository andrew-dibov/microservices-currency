package configs

import "time"

type AppConfig struct {
	Port string
	Prod bool

	Api   string
	Token string

	Cert string
	Key  string

	Services Services
	Infra    Infra

	Timeouts Timeouts
	Limits   Limits
}

type Services struct {
	Hist string
	Curr string
	Conv string
}

type Infra struct {
	Psql string

	PsqlDbMaxOpen     int
	PsqlDbMaxIdle     int
	PsqlDbMaxLifetime time.Duration
}

type Timeouts struct {
	Ex time.Duration

	Hist time.Duration
	Curr time.Duration
	Conv time.Duration
	Psql time.Duration

	Read     time.Duration
	Idle     time.Duration
	Write    time.Duration
	Shutdown time.Duration
}

type Limits struct {
}
