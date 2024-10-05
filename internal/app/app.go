package app

import (
	"context"
	"learn-golang-solid/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/samber/lo"
)

type App struct {
	db *sqlx.DB
	echo *echo.Echo
	cfg config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	return &App{
		db: lo.Must(datasource.NewDatabase(cfg.Database)),
		echo: middleware
	}
}
