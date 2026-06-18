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
	err := CheckNewWish(&newWish)
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
func CheckNewWish(newWish *model.NewWish) error {
	newWish.Title = strings.TrimSpace(newWish.Title)
	if newWish.Title == "" {
		return ErrorTitleEmpty
	}

	return nil
}
