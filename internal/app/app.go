package app

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"database/sql"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/api"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/store"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/middleware"
	"github.com/HySkySFC/hyskysfc-backend-api/migrations"
)

type Application struct {
	Logger *log.Logger
	PLTDHandler *api.PLTDHandler
	UserHandler *api.UserHandler
	TokenHandler *api.TokenHandler
	Middleware middleware.UserMiddleware
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open();
	if err != nil {
		return nil, err
	}
	
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pltdStore := store.NewPostgresPLTDStore(pgDB)
	userStore := store.NewPostgresUserStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)

	pltdHandler := api.NewPLTDHandler(pltdStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}

	app := &Application{
		Logger: logger,
		PLTDHandler: pltdHandler,
		UserHandler: userHandler,
		TokenHandler: tokenHandler,
		Middleware: middlewareHandler,
		DB: pgDB,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server status is available")
}
