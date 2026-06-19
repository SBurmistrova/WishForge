package handlers

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func HandlerWish(w http.ResponseWriter, r *http.Request) {
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
func HandlerWishID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/wishes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

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

func HandlerGetAddStep(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/wishes/")
	idStr = strings.TrimSuffix(idStr, "/steps")
	idWish, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	switch r.Method {
	case http.MethodGet:
		steps, err := service.GetSteps(idWish)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(w).Encode(steps)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case http.MethodPost:
		var newStep model.NewStep
		err := json.NewDecoder(r.Body).Decode(&newStep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newStep.IDWish = idWish

		step, err := service.CreateStep(newStep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(w).Encode(step)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	default:
		http.Error(w, "Unsupported request type", http.StatusMethodNotAllowed)
	}
}
func HandlerUpdateDeleteStep(w http.ResponseWriter, r *http.Request) {

	idWishStr1 := strings.TrimPrefix(r.URL.Path, "/wishes/")
	var idWishStr string
	for _, ch := range idWishStr1 {
		if ch == '/' {
			break
		}
		idWishStr = idWishStr + string(ch)
	}
	idWish, err := strconv.Atoi(idWishStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	idStepStr := strings.TrimPrefix(r.URL.Path, "/wishes/"+idWishStr+"/steps/")
	idStep, err := strconv.Atoi(idStepStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	switch r.Method {
	case http.MethodPatch:
		var updateStep model.Step
		err := json.NewDecoder(r.Body).Decode(&updateStep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		updateStep.ID = idStep
		updateStep.IDWish = idWish

		step, err := service.UpdateStep(updateStep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = json.NewEncoder(w).Encode(step)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

	case http.MethodDelete:
		err := service.DeleteStep(idWish, idStep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	default:
		http.Error(w, "Unsupported request type", http.StatusMethodNotAllowed)
	}
}
