package repositories

import "database/sql"

type PostgresRepository struct {
	postgresDatabase *sql.DB
}
