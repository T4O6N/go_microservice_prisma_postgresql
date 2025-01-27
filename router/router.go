package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tonpcst/go-microservice-prisma-postgresql/controllers"
)

func Routers() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Create /api prefix
	router.Route("/api", func(r chi.Router) {
		// Swagger docs
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:7000/api/swagger/doc.json"),
		))

		// User routes
		r.Post("/users", controllers.CreateUser)
		r.Get("/users/all", controllers.GetAllUsers)
		r.Get("/users/byId/{id}", controllers.GetUserByID)
		r.Put("/users/update/{id}", controllers.UpdateUser)
		r.Delete("/users/delete/{id}", controllers.DeleteUser)
	})

	return router
}
