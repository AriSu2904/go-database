package main

import (
	"context"
	"fmt"
	"go-db/config"
	"os"
)

func main() {

	dbConfig := config.GetConnection()
	defer dbConfig.Close()
	//

	//description := "Belajar Go - Database"
	//
	//query := "INSERT INTO m_todos (id, title, description) VALUES($1, $2, $3)"
	//
	//_, err := dbConfig.Exec(context.Background(), query, id, title, description)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//fmt.Println("Successfully add data todo!")

	id := "2"
	title := "Belajar"
	query := "SELECT description FROM m_todos WHERE id = $1 AND title = $2 LIMIT 1"

	var destination string
	err := dbConfig.QueryRow(context.Background(), query, id, title).Scan(&destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(destination)
}
