package handlers

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandlerWish(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		HandlerWishID(id, w, r)

		return
	}

	switch r.Method {
	case http.MethodGet:
		wishes, err := service.GetWishes()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(w).Encode(wishes)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	case http.MethodPost:
		var newWish model.NewWish
		err := json.NewDecoder(r.Body).Decode(&newWish)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = service.CreateWish(newWish)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	default:
		http.Error(w, "Unsupported request type", http.StatusMethodNotAllowed)
	}
}

func HandlerWishID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		wish, err := service.GetWish(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(w).Encode(wish)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	case http.MethodDelete:
		err := service.DeleteWish(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	default:
		http.Error(w, "Unsupported request type", http.StatusMethodNotAllowed)
	}
}
