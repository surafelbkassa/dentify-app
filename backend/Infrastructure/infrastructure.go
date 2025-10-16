package infrastructure

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
