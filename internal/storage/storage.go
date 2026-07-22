package storage

import (
	"WishForge/internal/model"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func GetWishes() ([]model.Wish, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return nil, err
	}
	defer dataBase.Close()

	request := "SELECT id, title, completed FROM wishes ORDER BY id"

	rows, err := dataBase.Query(request)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
func CreateWish(newWish model.CreateWishRequest) (int, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return 0, err
	}
	defer dataBase.Close()

	request := "INSERT INTO wishes (title) OUTPUT INSERTED.id VALUES (@title)"

	var id int
	err = dataBase.QueryRow(request, sql.Named("title", newWish.Title)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func GetWish(id int) (model.Wish, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return model.Wish{}, err
	}
	defer dataBase.Close()

	request := "SELECT id, title, completed FROM wishes WHERE id = @id"

	wish := model.Wish{}
	err = dataBase.QueryRow(request, sql.Named("id", id)).Scan(&wish.ID, &wish.Title, &wish.Completed)
	if err != nil {
		return model.Wish{}, err
	}

	return wish, nil
}
func UpdateWish(updateWish model.Wish) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := `DECLARE @old_title NVARCHAR(50)
                SET @old_title = (SELECT title FROM wishes WHERE id = @id)

                UPDATE wishes 
                SET title = COALESCE(NULLIF(@title, ''), @old_title), completed = @completed
                WHERE id = @id`

	_, err = dataBase.Exec(request,
		sql.Named("title", updateWish.Title),
		sql.Named("completed", updateWish.Completed),
		sql.Named("id", updateWish.ID))

	if err != nil {
		return err
	}
	return nil
}
func DeleteWish(id int) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	tx, err := dataBase.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.Exec("DELETE FROM steps WHERE id_wish = @id", sql.Named("id", id)); err != nil {
		return err
	}

	if _, err = tx.Exec("DELETE FROM wishes WHERE id = @id", sql.Named("id", id)); err != nil {
		return err
	}

	return tx.Commit()
}

func GetSteps(idWish int) ([]model.Step, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return nil, err
	}
	defer dataBase.Close()

	request := "SELECT id_wish, id, title, completed FROM steps WHERE id_wish = @id_wish ORDER BY id"

	rows, err := dataBase.Query(request, sql.Named("id_wish", idWish))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	steps := make([]model.Step, 0)

	for rows.Next() {
		var step model.Step

		err := rows.Scan(&step.IDWish, &step.ID, &step.Title, &step.Completed)
		if err != nil {
			return nil, err
		}

		steps = append(steps, step)
	}

	return steps, nil
}
func CreateStep(newStep model.CreateStep) (model.Step, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return model.Step{}, err
	}
	defer dataBase.Close()

	request := `DECLARE @id INT =
                COALESCE(
                ( SELECT MAX(id) + 1
                  FROM steps WITH (UPDLOCK)
                  WHERE id_wish = @id_wish ),
                1)

                INSERT INTO steps (id, id_wish, title) VALUES (@id, @id_wish, @title)
				
				SELECT @id`

	var id int
	err = dataBase.QueryRow(request,
		sql.Named("id_wish", newStep.IDWish),
		sql.Named("title", newStep.Title)).Scan(&id)

	if err != nil {
		return model.Step{}, err
	}

	return model.Step{IDWish: newStep.IDWish, ID: id, Title: newStep.Title, Completed: false}, nil
}
func UpdateStep(updateStep model.Step) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := `UPDATE steps 
                SET title = @title, completed = @completed
                WHERE id_wish = @id_wish AND id = @id`

	_, err = dataBase.Exec(request,
		sql.Named("title", updateStep.Title),
		sql.Named("completed", updateStep.Completed),
		sql.Named("id_wish", updateStep.IDWish),
		sql.Named("id", updateStep.ID))

	if err != nil {
		return err
	}
	return nil
}
func DeleteStep(idWish int, idStep int) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := `DELETE FROM steps WHERE id_wish = @id_wish AND id = @id
	            UPDATE steps SET id = id - 1 WHERE id_wish = @id_wish AND id > @id`

	_, err = dataBase.Exec(request,
		sql.Named("id_wish", idWish),
		sql.Named("id", idStep))

	if err != nil {
		return err
	}

	return nil
}

func connectDataBase() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %s", err.Error())
	}

	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("sqlserver://@%s:%s?database=%s", server, port, database)

	dataBase, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("Error creating connection: %v", err)
	}

	err = dataBase.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	return dataBase, nil
}
