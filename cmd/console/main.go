package main

import (
	"context"
	"database/sql"
	"errhello/model"
	"errhello/service/impl"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "zenith"
	dbname   = "rant"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	userService := impl.New(db)
	user := model.User{}
	err = userService.CreateUser(ctx, &user)
	if err != nil {
		log.Printf(err.Error())
	}
}
