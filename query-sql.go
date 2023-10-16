package main

import (
	"context"
	"fmt"
	"go-db/config"
)

func main() {

	database := config.GetConnection()
	defer database.Close()
	query := "SELECT * FROM m_todos"

	rows, err := database.Query(context.Background(), query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		//Data Penampung
		var id, title, description string

		err := rows.Scan(&id, &title, &description)

		if err != nil {
			panic(err)
		}
		fmt.Println(id, title, description)
	}

}
