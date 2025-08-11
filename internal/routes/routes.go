package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/app"
)	

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func (r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		// PLTD
		r.Get("/pltd", app.Middleware.RequireUser(app.PLTDHandler.HandleGetAllPLTD))
		r.Get("/pltd/{id}", app.Middleware.RequireUser(app.PLTDHandler.HandleGetPLTDByID))
		r.Post("/pltd", app.Middleware.RequireUser(app.PLTDHandler.HandleCreatePLTD))
		r.Put("/pltd/{id}", app.Middleware.RequireUser(app.PLTDHandler.HandleUpdatePLTDByID))
		r.Delete("/pltd/{id}", app.Middleware.RequireUser(app.PLTDHandler.HandleDeletePLTDByID))

		// PLTS
		r.Post("/plts", app.Middleware.RequireUser(app.PLTSHandler.HandleReplaceAllPLTS))
		r.Get("/plts", app.Middleware.RequireUser(app.PLTSHandler.HandleGetAllPLTS))
	})

	r.Get("/health", app.HealthCheck)

	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/token/authentication", app.TokenHandler.HandleCreateToken)

	return r
}

