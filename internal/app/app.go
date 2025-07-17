package app

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/api"
)

type Application struct {
	Logger *log.Logger
	PLTDHandler *api.PLTDHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pltdHandler := api.NewPLTDHandler()
	app := &Application{
		Logger: logger,
		PLTDHandler: pltdHandler,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server status is available")
}
