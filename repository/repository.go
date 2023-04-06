package repository

import (
	"database/sql"
	"mopoen-remake/repository/postgres"

	_ "github.com/lib/pq"
)

type Repository struct {
	SQLConn *sql.DB
	*postgres.Queries
}

func NewRepository(driver string, source string) (Repository, error) {
	SQLConn, err := sql.Open(driver, source)
	if err != nil {
		return Repository{}, err
	}

	repo := Repository{
		SQLConn: SQLConn,
		Queries: postgres.New(SQLConn),
	}

	return repo, nil
}
