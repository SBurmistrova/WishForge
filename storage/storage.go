package storage

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type Wish struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var DataBase *sql.DB

func GetWishes() ([]Wish, error) {
	request := "SELECT ID, name, completed FROM wishes"

	rows, err := DataBase.Query(request)
	if err != nil {
		return nil, err
	}

	wishes := make([]Wish, 0)

	for rows.Next() {
		var wish Wish

		err := rows.Scan(&wish.ID, &wish.Name, &wish.Completed)
		if err != nil {
			return nil, err
		}

		wishes = append(wishes, wish)
	}

	return wishes, nil
}
