package storage

import (
	"WishForge/internal/model"
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var DataBase *sql.DB

func GetWishes() ([]model.Wish, error) {
	request := "SELECT ID, name, completed FROM wishes"

	rows, err := DataBase.Query(request)
	if err != nil {
		return nil, err
	}

	wishes := make([]model.Wish, 0)

	for rows.Next() {
		var wish model.Wish

		err := rows.Scan(&wish.ID, &wish.Name, &wish.Completed)
		if err != nil {
			return nil, err
		}

		wishes = append(wishes, wish)
	}

	return wishes, nil
}
