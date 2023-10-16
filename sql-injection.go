package main

import (
	"context"
	"fmt"
	"go-db/config"
)

func main() {

	database := config.GetConnection()
	defer database.Close()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM account WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"

	var resultUsername string
	err := database.QueryRow(context.Background(), script).Scan(resultUsername)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success login")
}
