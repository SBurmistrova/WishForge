package handlers

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// GetWishes godoc
//
// @Summary      Get all wishes
// @Tags         wishes
// @Produce      json
// @Success      200 {array} model.Wish
// @Failure      500 {object} model.Error
// @Router       /wishes [get]
func GetWishes(w http.ResponseWriter, r *http.Request) {
	wishes, err := service.GetWishes()

	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)

		return
	}

	sendJSON(w, wishes)
}

// PostWish godoc
//
// @Summary      Create a new wish
// @Tags         wishes
// @Accept       json
// @Produce      json
// @Param        wish body model.CreateWishRequest true "Wish"
// @Success      200 {object} model.Wish
// @Failure      400 {object} model.Error
// @Router       /wishes [post]
func PostWish(w http.ResponseWriter, r *http.Request) {
	var cw model.CreateWishRequest
	err := json.NewDecoder(r.Body).Decode(&cw)

	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	wish, err := service.CreateWish(cw)

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

// GetWish godoc
//
// @Summary      Get wish
// @Tags         wishes
// @Produce      json
// @Param        wishID path int true "Wish ID"
// @Success      200 {object} model.Wish
// @Failure      404 {object} model.Error
// @Router       /wishes/{wishID} [get]
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

// DeleteWish godoc
//
// @Summary      Delete a wish
// @Tags         wishes
// @Param        wishID path int true "Wish ID"
// @Success      200
// @Failure      404 {object} model.Error
// @Router       /wishes/{wishID} [delete]
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

// PatchWish godoc
//
// @Summary      Change of wish
// @Tags         wishes
// @Accept       json
// @Produce      json
// @Param        wishID path int true "Wish ID"
// @Param        wish body model.UpdateWishRequest true "Wish"
// @Success      200 {object} model.Wish
// @Failure      400 {object} model.Error
// @Router       /wishes/{wishID} [patch]
func PatchWish(w http.ResponseWriter, r *http.Request) {
	var uwr model.UpdateWishRequest

	err := json.NewDecoder(r.Body).Decode(&uwr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	ID, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	updateWish := model.Wish{
		ID:        ID,
		Title:     uwr.Title,
		Completed: uwr.Completed,
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

// GetSteps godoc
//
// @Summary      Getting all the steps for a wish
// @Tags         steps
// @Param        wishID path int true "Wish ID"
// @Produce      json
// @Success      200 {array} model.Step
// @Failure      500 {object} model.Error
// @Router       /wishes/{wishID}/steps [get]
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

// PostStep godoc
//
// @Summary      Create a new step
// @Tags         steps
// @Accept       json
// @Produce      json
// @Param        wishID path int true "Wish ID"
// @Param        wish body model.CreateStepRequest true "Step"
// @Success      200 {object} model.Step
// @Failure      400 {object} model.Error
// @Router       /wishes/{wishID}/steps [post]
func PostStep(w http.ResponseWriter, r *http.Request) {
	var csr model.CreateStepRequest
	err := json.NewDecoder(r.Body).Decode(&csr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	IDWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	newStep := model.CreateStep{
		IDWish: IDWish,
		Title:  csr.Title,
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

// PatchStep godoc
//
// @Summary      Change of step
// @Tags         steps
// @Accept       json
// @Produce      json
// @Param        wishID path int true "Wish ID"
// @Param        stepID path int true "Step ID"
// @Param        wish body model.UpdateStepRequest true "Step"
// @Success      200 {object} model.Step
// @Failure      400 {object} model.Error
// @Router       /wishes/{wishID}/steps/{stepID} [patch]
func PatchStep(w http.ResponseWriter, r *http.Request) {
	var usr model.UpdateStepRequest
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	IDWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	ID, err := getID(r, "stepID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	updateStep := model.Step{
		IDWish:    IDWish,
		ID:        ID,
		Title:     usr.Title,
		Completed: usr.Completed,
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

// DeleteStep godoc
//
// @Summary      Delete a step
// @Tags         steps
// @Param        wishID path int true "Wish ID"
// @Param        stepID path int true "Step ID"
// @Success      200
// @Failure      404 {object} model.Error
// @Router       /wishes/{wishID}/steps/{stepID} [delete]
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

// GetProgress godoc
//
// @Summary      Get progress
// @Tags         progress
// @Produce      json
// @Param        wishID path int true "Wish ID"
// @Success      200 {object} model.Progress
// @Failure      404 {object} model.Error
// @Router       /wishes/{wishID}/progress [get]
func Progress(w http.ResponseWriter, r *http.Request) {
	idWish, err := getID(r, "wishID")
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusBadRequest)

		return
	}

	progress, err := service.GetProgress(idWish)
	if err != nil {
		sendError(w, model.Error{Text: err.Error()}, http.StatusInternalServerError)

		return
	}

	sendJSON(w, progress)
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
