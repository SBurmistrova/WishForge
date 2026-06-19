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

	request := "SELECT id, title, completed FROM wishes ORDER BY id"

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
func CreateWish(newWish model.NewWish) (int, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return 0, err
	}
	defer dataBase.Close()

	request := "INSERT INTO wishes (title) OUTPUT INSERTED.id VALUES (@title)"

	rows, err := dataBase.Query(request, sql.Named("title", newWish.Title))
	if err != nil {
		return 0, err
	}

	var id int

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
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
func UpdateWish(updateWish model.Wish) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := `UPDATE wishes 
                SET title = @title, completed = @completed
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

	request := `DELETE FROM steps WHERE id_wish = @id
	            DELETE FROM wishes WHERE id = @id`

	_, err = dataBase.Exec(request, sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
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
func CreateStep(newStep model.NewStep) (model.Step, error) {
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

	rows, err := dataBase.Query(request,
		sql.Named("id_wish", newStep.IDWish),
		sql.Named("title", newStep.Title))

	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return model.Step{}, err
		}
	}

	if err != nil {
		return model.Step{}, err
	}

	return model.Step{IDWish: newStep.IDWish, ID: id, Title: newStep.Title, Completed: false}, nil
}
func UpdateStep(updateStep model.Step) (model.Step, error) {
	dataBase, err := connectDataBase()
	if err != nil {
		return model.Step{}, err
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
		return model.Step{}, err
	}
	return updateStep, nil
}
func DeleteStep(idWish int, idStep int) error {
	dataBase, err := connectDataBase()
	if err != nil {
		return err
	}
	defer dataBase.Close()

	request := "DELETE FROM steps WHERE id_wish = @id_wish AND id = @id"

	_, err = dataBase.Exec(request,
		sql.Named("id_wish", idWish),
		sql.Named("id", idStep))

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
