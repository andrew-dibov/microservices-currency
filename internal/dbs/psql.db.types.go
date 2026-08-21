package dbs

import "database/sql"

type PsqlDB struct {
	*sql.DB
}
