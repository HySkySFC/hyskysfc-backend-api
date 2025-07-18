package app

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"database/sql"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/api"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/store"
	"github.com/HySkySFC/hyskysfc-backend-api/migrations"
)

type Application struct {
	Logger *log.Logger
	PLTDHandler *api.PLTDHandler
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

	pltdHandler := api.NewPLTDHandler(pltdStore, logger)
	app := &Application{
		Logger: logger,
		PLTDHandler: pltdHandler,
		DB: pgDB,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server status is available")
}
