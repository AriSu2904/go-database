package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {

	dbUrl := "postgres://postgres:admin@127.0.0.1/latihan"

	connect, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Println("Unable to connect to database :", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Successfully connected!")
	}
	defer connect.Close(context.Background())

}
