package apiserver

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/pitshifer/valera-acceptor/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	s := NewServer(store)
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
