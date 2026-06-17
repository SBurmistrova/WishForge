package main

import (
	"WishForge/internal/handlers"
	"WishForge/internal/storage"
	"database/sql"
	"fmt"
	"net/http"
)

func main() {
	connString := "server=localhost;database=ForLearning;integrated security=true"
	dataBase, err := sql.Open("sqlserver", connString)

	if err != nil {
		fmt.Println(err)
	}

	if err := dataBase.Ping(); err != nil {
		fmt.Println(err)
	}

	defer dataBase.Close()

	storage.DataBase = dataBase

	http.HandleFunc("/", handlers.GetWishes)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
