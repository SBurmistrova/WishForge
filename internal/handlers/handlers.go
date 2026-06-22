package handlers

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetWishes(w http.ResponseWriter, r *http.Request) {
	wishes, err := service.GetWishes()

	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wishes)
}
func PostWish(w http.ResponseWriter, r *http.Request) {
	var newWish model.NewWish
	err := json.NewDecoder(r.Body).Decode(&newWish)

	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	wish, err := service.CreateWish(newWish)

	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wish)
}
func GetWish(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "wishID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	wish, err := service.GetWish(id)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wish)
}
func DeleteWish(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "wishID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	err = service.DeleteWish(id)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}
}
func PatchWish(w http.ResponseWriter, r *http.Request) {
	var updateWish model.Wish

	err := json.NewDecoder(r.Body).Decode(&updateWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	idStr := chi.URLParam(r, "wishID")
	updateWish.ID, err = strconv.Atoi(idStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	wish, err := service.UpdateWish(updateWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wish)
}
func GetStaps(w http.ResponseWriter, r *http.Request) {
	idWishStr := chi.URLParam(r, "wishID")
	idWish, err := strconv.Atoi(idWishStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	steps, err := service.GetSteps(idWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}
func PostStep(w http.ResponseWriter, r *http.Request) {
	var newStep model.NewStep
	err := json.NewDecoder(r.Body).Decode(&newStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
		return
	}

	idWishStr := chi.URLParam(r, "wishID")
	newStep.IDWish, err = strconv.Atoi(idWishStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	step, err := service.CreateStep(newStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(step)
}
func PatchStep(w http.ResponseWriter, r *http.Request) {
	var updateStep model.Step
	err := json.NewDecoder(r.Body).Decode(&updateStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	idWishStr := chi.URLParam(r, "wishID")
	updateStep.IDWish, err = strconv.Atoi(idWishStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	idStepStr := chi.URLParam(r, "stepID")
	updateStep.ID, err = strconv.Atoi(idStepStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	step, err := service.UpdateStep(updateStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(step)
}
func DeleteStep(w http.ResponseWriter, r *http.Request) {
	idWishStr := chi.URLParam(r, "wishID")
	idWish, err := strconv.Atoi(idWishStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	idStepStr := chi.URLParam(r, "stepID")
	idStep, err := strconv.Atoi(idStepStr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}

	err = service.DeleteStep(idWish, idStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
	}
}

func sendError(w http.ResponseWriter, err model.Error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
