package main

import (
	"WishForge/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	r.Route("/wishes", func(r chi.Router) {
		r.Get("/", handlers.GetWishes)
		r.Post("/", handlers.PostWish)

		r.Route("/{wishID}", func(r chi.Router) {
			r.Get("/", handlers.GetWish)
			r.Delete("/", handlers.DeleteWish)
			r.Patch("/", handlers.PatchWish)

			r.Route("/steps", func(r chi.Router) {
				r.Get("/", handlers.GetStaps)
				r.Post("/", handlers.PostStep)

				r.Route("/{stepID}", func(r chi.Router) {
					r.Patch("/", handlers.PatchStep)
					r.Delete("/", handlers.DeleteStep)
				})
			})
		})
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println(err)
	}
}
