package repo_test

import (
	"os"
	"testing"

	"github.com/forroots/nomity-admin-api-v3/internal/infra/db"
	appdb "github.com/forroots/nomity-admin-api-v3/internal/infra/db"
	"github.com/uptrace/bun"
)

var testDB *bun.DB

func TestMain(m *testing.M) {
	params := db.DBParams{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5430,
		User:     "nomity_dev",
		Password: "nomity_dev",
		DBName:   "nomity_dev",
		SSLMode:  "disable",
	}

	var err error
	testDB, err = appdb.NewBunDB(params, true)
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	code := m.Run()

	_ = testDB.Close()
	os.Exit(code)
}
