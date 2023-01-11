package dbrepo

import (
	"database/sql"

	"github.com/cxt314/learning-bnb-webapp/internal/config"
	"github.com/cxt314/learning-bnb-webapp/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
