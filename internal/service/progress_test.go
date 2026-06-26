package service_test

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestGetProgress(t *testing.T) {

	tests := []struct {
		name             string
		steps            []string
		expectedProgress model.Progress
	}{
		{"Returns100PercentWhenAllStepsCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 1)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 1)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 3, 'step3', 1)`},
			model.Progress{Progress: 100, CountCompleted: 3, CountNotCompleted: 0}},
		{"Returns50PercentWhenHalfStepsCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 1)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 1)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 3, 'step3', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 4, 'step4', 0)`},
			model.Progress{Progress: 50, CountCompleted: 2, CountNotCompleted: 2}},
		{"Returns33PercentWhenOneOfThreeStepsCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 1)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 3, 'step3', 0)`},
			model.Progress{Progress: 33, CountCompleted: 1, CountNotCompleted: 2}},
		{"ReturnsZeroWhenNoStepsCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 3, 'step3', 0)`},
			model.Progress{Progress: 0, CountCompleted: 0, CountNotCompleted: 3}},
		{"ReturnsEmptyProgressWhenWishHasNoSteps", []string{}, model.Progress{}},
	}

	db, err := connectDataBase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db.Exec("DELETE FROM steps")
			db.Exec("DELETE FROM wishes")

			var wishID int

			err := db.QueryRow(`INSERT INTO wishes (title) OUTPUT INSERTED.id VALUES ('test')`).Scan(&wishID)
			if err != nil {
				t.Fatal(err)
			}

			for _, step := range test.steps {
				db.Exec(step, sql.Named("id_wish", wishID))
			}

			result, err := service.GetProgress(int(wishID))
			if err != nil {
				t.Fatal(err)
			}

			if result.CountCompleted != test.expectedProgress.CountCompleted {
				t.Errorf("Expected %d completed, got %d", test.expectedProgress.CountCompleted, result.CountCompleted)
			}

			if result.CountNotCompleted != test.expectedProgress.CountNotCompleted {
				t.Errorf("Expected %d not completed, got %d", test.expectedProgress.CountNotCompleted, result.CountNotCompleted)
			}

			if result.Progress != test.expectedProgress.Progress {
				t.Errorf("Expected %d%%, got %d", test.expectedProgress.Progress, result.Progress)
			}

			defer db.Exec("DELETE FROM steps WHERE id_wish = @id", sql.Named("id", wishID))
			defer db.Exec("DELETE FROM wishes WHERE id = @id", sql.Named("id", wishID))
		})
	}
}

func connectDataBase() (*sql.DB, error) {
	server := "localhost"
	port := "1433"
	database := "ForLearning_test"

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
