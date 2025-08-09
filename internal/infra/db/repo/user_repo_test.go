package repo_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/forroots/nomity-admin-api-v3/internal/infra/db"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/db/repo"
)

func TestUserRepo(t *testing.T) {
	params := db.DBParams{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5430,
		User:     "nomity_dev",
		Password: "nomity_dev",
		DBName:   "nomity_dev",
		SSLMode:  "disable",
	}

	bdb, err := db.NewBunDB(params, true)
	if err != nil {
		log.Fatal(err)
	}
	defer bdb.Close()

	userRepo := repo.NewUserRepo(bdb)

	// 全件取得
	users, err := userRepo.ListAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users:", users)

	// 単一取得
	u, err := userRepo.FindByID(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User 1:", u)
}
