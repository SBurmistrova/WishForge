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
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)

		return
	}

	sendJSON(w, wishes)
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
		if err == service.ErrorTitleEmpty {
			sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
		} else {
			sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
		}

		return
	}

	sendJSON(w, wish)
}
func GetWish(w http.ResponseWriter, r *http.Request) {
	idWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	wish, err := service.GetWish(idWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)

		return
	}

	sendJSON(w, wish)
}
func DeleteWish(w http.ResponseWriter, r *http.Request) {
	idWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	err = service.DeleteWish(idWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
	}
}
func PatchWish(w http.ResponseWriter, r *http.Request) {
	var updateWish model.Wish

	err := json.NewDecoder(r.Body).Decode(&updateWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	updateWish.ID, err = getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	wish, err := service.UpdateWish(updateWish)
	if err != nil {
		if err == service.ErrorTitleEmpty {
			sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
		} else {
			sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
		}

		return
	}

	sendJSON(w, wish)
}
func GetSteps(w http.ResponseWriter, r *http.Request) {
	idWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	steps, err := service.GetSteps(idWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)

		return
	}

	sendJSON(w, steps)
}
func PostStep(w http.ResponseWriter, r *http.Request) {
	var newStep model.NewStep
	err := json.NewDecoder(r.Body).Decode(&newStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	newStep.IDWish, err = getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	step, err := service.CreateStep(newStep)
	if err != nil {
		if err == service.ErrorTitleEmpty {
			sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
		} else {
			sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
		}

		return
	}

	sendJSON(w, step)
}
func PatchStep(w http.ResponseWriter, r *http.Request) {
	var updateStep model.Step
	err := json.NewDecoder(r.Body).Decode(&updateStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	updateStep.IDWish, err = getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	updateStep.ID, err = getID(r, "stepID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	step, err := service.UpdateStep(updateStep)
	if err != nil {
		if err == service.ErrorTitleEmpty {
			sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)
		} else {
			sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
		}

		return
	}

	sendJSON(w, step)
}
func DeleteStep(w http.ResponseWriter, r *http.Request) {
	idWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	idStep, err := getID(r, "stepID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	err = service.DeleteStep(idWish, idStep)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)
	}
}

func sendError(w http.ResponseWriter, err model.Error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
func sendJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func getID(r *http.Request, key string) (int, error) {
	idStr := chi.URLParam(r, key)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
