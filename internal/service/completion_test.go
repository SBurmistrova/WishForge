package service_test

import (
	"WishForge/internal/model"
	"WishForge/internal/service"
	"database/sql"
	"testing"
)

func TestAutoCompleteWish(t *testing.T) {
	tests := []struct {
		name              string
		steps             []string
		expectedStateWish bool
	}{
		{"CompletesWishWhenLastStepBecomesCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 0)`},
			true},
		{"CompletesWishWhenAllStepsBecomeCompleted", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 1)`},
			true},
		{"KeepsWishIncompleteWhenRemainingStepsExist", []string{
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 1, 'step1', 0)`,
			`INSERT INTO steps (id_wish, id, title, completed)
             VALUES (@id_wish, 2, 'step2', 0)`},
			false},
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

			_, err = service.UpdateStep(model.Step{IDWish: int(wishID), ID: 1, Title: "step1", Completed: true})
			if err != nil {
				t.Fatal(err)
			}

			wish, err := service.GetWish(int(wishID))
			if err != nil {
				t.Fatal(err)
			}

			if wish.Completed != test.expectedStateWish {
				t.Errorf("Expected state wish: %v, got: %v", test.expectedStateWish, wish.Completed)
			}

			defer db.Exec("DELETE FROM steps WHERE id_wish = @id", sql.Named("id", wishID))
			defer db.Exec("DELETE FROM wishes WHERE id = @id", sql.Named("id", wishID))
		})
	}
}
