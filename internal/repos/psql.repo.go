package repos

import (
	"database/sql"
)

type PsqlRepo struct {
	psql *sql.DB
}

func NewPsqlRepo(psql *sql.DB) *PsqlRepo {
	return &PsqlRepo{
		psql: psql,
	}
}
