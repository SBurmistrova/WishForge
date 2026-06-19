package service

import (
	"WishForge/internal/model"
	"WishForge/internal/storage"
	"errors"
	"strings"
)

var ErrorTitleEmpty = errors.New("Title is empty")
var ErrorWishNotFound = errors.New("Wish is not found")

func GetWishes() ([]model.Wish, error) {
	return storage.GetWishes()
}
func CreateWish(newWish model.NewWish) error {
	err := CheckTitle(&newWish.Title)
	if err != nil {
		return err
	}

	err = storage.CreateWish(newWish)
	if err != nil {
		return err
	}

	return nil
}
func GetWish(id int) (model.Wish, error) {
	wish, err := storage.GetWish(id)
	if err != nil {
		return model.Wish{}, err
	}

	if wish.ID == 0 {
		return model.Wish{}, ErrorWishNotFound
	}

	return wish, nil
}
func DeleteWish(id int) error {
	err := storage.DeleteWish(id)
	if err != nil {
		return err
	}

	return nil
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

	step, err := storage.UpdateStep(updateStep)
	if err != nil {
		return model.Step{}, err
	}

	return step, nil
}
func DeleteStep(idWish int, idStep int) error {
	err := storage.DeleteStep(idWish, idStep)
	if err != nil {
		return err
	}

	return nil
}

func CheckTitle(title *string) error {
	*title = strings.TrimSpace(*title)
	if *title == "" {
		return ErrorTitleEmpty
	}

	return nil
}
