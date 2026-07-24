package main

import (
	"WishForge/internal/handlers"
	"WishForge/internal/middleware"
	"fmt"
	"net/http"

	_ "WishForge/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           WishForge API
// @version         1.0
// @description     REST API for managing wishes and steps.
// @BasePath        /
func main() {
	logger, err := middleware.CreateLogger()

	r := chi.NewRouter()

	r.Use(middleware.RequestLogger(logger))

	r.Get("/swagger/*", httpSwagger.Handler())

	r.Route("/wishes", func(r chi.Router) {
		r.Get("/", handlers.GetWishes)
		r.Post("/", handlers.PostWish)

		r.Route("/{wishID}", func(r chi.Router) {
			r.Get("/", handlers.GetWish)
			r.Delete("/", handlers.DeleteWish)
			r.Patch("/", handlers.PatchWish)
			r.Get("/progress", handlers.Progress)

			r.Route("/steps", func(r chi.Router) {
				r.Get("/", handlers.GetSteps)
				r.Post("/", handlers.PostStep)

				r.Route("/{stepID}", func(r chi.Router) {
					r.Patch("/", handlers.PatchStep)
					r.Delete("/", handlers.DeleteStep)
				})
			})
		})
	})

	err = http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println(err)
	}
}
