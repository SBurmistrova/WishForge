package storage

import (
	"WishForge/internal/model"
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetWishes() ([]model.Wish, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return nil, err
	}
	defer dataBase.Close()

	request := "SELECT id, title, completed FROM wishes"

	rows, err := dataBase.Query(request)
	if err != nil {
		return nil, err
	}

	wishes := make([]model.Wish, 0)

	for rows.Next() {
		var wish model.Wish

		err := rows.Scan(&wish.ID, &wish.Title, &wish.Completed)
		if err != nil {
			return nil, err
		}

		wishes = append(wishes, wish)
	}

	return wishes, nil
}
func CreateWish(newWish model.NewWish) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := "INSERT INTO wishes (title) VALUES (@title)"

	_, err = dataBase.Exec(request, sql.Named("title", newWish.Title))
	if err != nil {
		return err
	}

	return nil
}
func GetWish(id int) (model.Wish, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return model.Wish{}, err
	}
	defer dataBase.Close()

	request := "SELECT id, title, completed FROM wishes WHERE id = @id"

	rows, err := dataBase.Query(request, sql.Named("id", id))
	if err != nil {
		return model.Wish{}, err
	}

	wish := model.Wish{}
	for rows.Next() {
		err := rows.Scan(&wish.ID, &wish.Title, &wish.Completed)
		if err != nil {
			return model.Wish{}, err
		}
	}

	return wish, nil
}
func DeleteWish(id int) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := "DELETE FROM wishes WHERE id = @id"

	_, err = dataBase.Exec(request, sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}

func connectDataBase() (*sql.DB, error) {
	connString := "server=localhost;database=ForLearning;integrated security=true"
	dataBase, err := sql.Open("sqlserver", connString)

	if err != nil {
		return nil, err
	}

	if err := dataBase.Ping(); err != nil {
		return nil, err
	}

	return dataBase, nil
}
