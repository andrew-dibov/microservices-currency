package dbs

import "database/sql"

type PsqlDb struct {
	*sql.DB
}
