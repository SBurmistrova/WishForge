package main

import (
	"WishForge/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/wishes", handlers.HandlerWish)
	http.HandleFunc("/wishes/{id}", handlers.HandlerWishID)
	http.HandleFunc("/wishes/{wishID}/steps", handlers.HandlerGetAddStep)
	http.HandleFunc("/wishes/{wishID}/steps/{stepID}", handlers.HandlerUpdateDeleteStep)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
