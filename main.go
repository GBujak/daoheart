package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/gbujak/daoheart-go/m/v2/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	goose.SetBaseFS(embedMigrations)
	err := goose.SetDialect("sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	defer ctx.Done()

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln(err)
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.New(db)

	users, err := repo.FindAllUsers(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for idx, user := range users {
		fmt.Printf("-> User no. %v with name %v\n", idx, user.Username)
	}
}
