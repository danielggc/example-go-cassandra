package routes

import (
	"broker/cmd/api/infrastruture/cassandra"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Config struct {
	UserRepository cassandra.UserRepository
}

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "content-type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300,
			},
		),
	)

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Post("/", app.Broker)
	mux.Get("/getUserById/{id}", app.GetUserById)
	mux.Post("/SaveUserById", app.SaveUser)

	return mux

}
