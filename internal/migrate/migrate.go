package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pitshifer/valera-acceptor/internal/app/apiserver"
)

// Do applies migrations
func Do(config *apiserver.Config, action string) error {
	m, err := migrate.New("file://migrations", config.DatabaseURL)
	if err != nil {
		return err
	}
	if action == "up" {
		if err = m.Up(); err != nil {
			return err
		}
		return nil
	}
	if action == "down" {
		if err = m.Down(); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("wrong action for migrate: action is \"%s\"", action)
}
