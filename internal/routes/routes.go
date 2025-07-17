package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/app"
)	

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/pltd/{id}", app.PLTDHandler.HandleGetPLTDByID)
	
	r.Post("/pltd", app.PLTDHandler.HandleCreatePLTD)
	return r
}

