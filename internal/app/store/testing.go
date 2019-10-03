package store

import (
	"fmt"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)); err != nil {
					t.Fatal(err)
				}
			}
		}

		s.Close()
	}
}
