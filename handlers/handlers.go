package handlers

import (
	"WishForge/storage"
	"encoding/json"
	"net/http"
)

func GetWishes(w http.ResponseWriter, r *http.Request) {
	wishes, err := storage.GetWishes()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	err = json.NewEncoder(w).Encode(wishes)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}
}
