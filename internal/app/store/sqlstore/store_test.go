package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "root:dfnheif@(localhost:5506)/acceptor_test?parseTime=true"
	}

	os.Exit(m.Run())
}
