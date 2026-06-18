package main

import (
	"WishForge/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/wishes", handlers.HandlerWish)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
