package sqlstore

import (
	"database/sql"
	"fmt"
	"testing"
)

// TestDB ...
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				if _, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)); err != nil {
					t.Fatal(err)
				}
			}
		}

		db.Close()
	}
}
