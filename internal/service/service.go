package service

import (
	"WishForge/internal/model"
	"WishForge/internal/storage"
	"errors"
	"strings"
)

var ErrorTitleEmpty = errors.New("Title is empty")

func GetWishes() ([]model.Wish, error) {
	return storage.GetWishes()
}
func CreateWish(newWish model.NewWish) (model.Wish, error) {
	err := CheckTitle(&newWish.Title)
	if err != nil {
		return model.Wish{}, err
	}

	id, err := storage.CreateWish(newWish)
	if err != nil {
		return model.Wish{}, err
	}

	return model.Wish{ID: id, Title: newWish.Title, Completed: false}, nil
}
func GetWish(id int) (model.Wish, error) {
	wish, err := storage.GetWish(id)
	if err != nil {
		return model.Wish{}, err
	}

	return wish, nil
}
func UpdateWish(updateWish model.Wish) (model.Wish, error) {
	if !updateWish.Completed {
		err := CheckTitle(&updateWish.Title)
		if err != nil {
			return model.Wish{}, err
		}
	}

	err := storage.UpdateWish(updateWish)
	if err != nil {
		return model.Wish{}, err
	}

	return model.Wish{ID: updateWish.ID, Title: updateWish.Title, Completed: updateWish.Completed}, nil
}
func DeleteWish(id int) error {
	return storage.DeleteWish(id)
}

func GetSteps(idWish int) ([]model.Step, error) {
	return storage.GetSteps(idWish)
}
func CreateStep(newStep model.NewStep) (model.Step, error) {
	err := CheckTitle(&newStep.Title)
	if err != nil {
		return model.Step{}, err
	}

	step, err := storage.CreateStep(newStep)
	if err != nil {
		return model.Step{}, err
	}

	return step, nil
}
func UpdateStep(updateStep model.Step) (model.Step, error) {
	err := CheckTitle(&updateStep.Title)
	if err != nil {
		return model.Step{}, err
	}

	err = storage.UpdateStep(updateStep)
	if err != nil {
		return model.Step{}, err
	}

	if updateStep.Completed {
		progress, err := storage.GetProgress(updateStep.IDWish)
		if err != nil {
			return model.Step{}, err
		}

		if progress.Progress == 100 {
			_, err := UpdateWish(model.Wish{ID: updateStep.IDWish, Completed: true})
			if err != nil {
				return model.Step{}, err
			}
		}
	}

	return updateStep, nil
}
func DeleteStep(idWish int, idStep int) error {
	return storage.DeleteStep(idWish, idStep)
}

func GetProgress(idWish int) (model.Progress, error) {
	return storage.GetProgress(idWish)
}

func CheckTitle(title *string) error {
	*title = strings.TrimSpace(*title)
	if *title == "" {
		return ErrorTitleEmpty
	}

	return nil
}
