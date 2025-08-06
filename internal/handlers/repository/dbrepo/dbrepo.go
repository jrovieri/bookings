package dbrepo

import (
	"database/sql"

	"github.com/jrovieri/bookings/internal/config"
	"github.com/jrovieri/bookings/internal/handlers/repository"
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
