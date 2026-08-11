package configs

import "time"

type AppConfig struct {
	Port string
	Prod bool

	Api   string
	Token string

	Services       Services
	Infrastructure Infrastructure

	Timeouts Timeouts
	Limits   Limits
}

type Services struct {
	Hist string
	Curr string
	Conv string
}

type Infrastructure struct {
	Psql string
}

type Timeouts struct {
	Ex time.Duration

	Hist time.Duration
	Curr time.Duration
	Conv time.Duration

	Read     time.Duration
	Idle     time.Duration
	Write    time.Duration
	Shutdown time.Duration
}

type Limits struct{}
