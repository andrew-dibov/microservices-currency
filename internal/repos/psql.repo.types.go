package repos

import (
	"database/sql"
)

type PsqlRepo struct {
	psql *sql.DB
}
