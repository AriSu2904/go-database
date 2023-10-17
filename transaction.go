package main

import (
	"context"
	"fmt"
	"go-db/config"
	"strconv"
)

func main() {

	database := config.GetConnection()
	defer database.Close()

	tx, err := database.Begin(context.Background())

	if err != nil {
		panic(err)
	}

	//transaction process
	query := "INSERT INTO m_todos (id, title, description) VALUES ($1, $2, $3)"
	for i := 11; i <= 20; i++ {
		id := strconv.Itoa(i)
		title := "Belajar Day-" + strconv.Itoa(i)
		description := "Belajar Golang Database Hari Ke-" + strconv.Itoa(i)

		_, err = tx.Exec(context.Background(), query, id, title, description)
		fmt.Println("Proses Ke-", i)
		if err != nil {
			panic(err)
		}
	}
	//Transaction Commit -> ganti jadi transaction Rollback
	err = tx.Rollback(context.Background())
	if err != nil {
		panic(err)
	}
}
