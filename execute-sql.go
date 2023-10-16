package main

import (
	"context"
	"fmt"
	"go-db/config"
)

func main() {

	// Panggil fungsi untuk koneksi ke database
	dbConfig := config.GetConnection()
	defer dbConfig.Close()

	query := "INSERT INTO m_todos (id, title, description) VALUES ('1', 'Belajar', 'Belajar Golang Dasar')"

	_, err := dbConfig.Exec(context.Background(), query)

	//throw atau panic saat terjadi error
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully add data todo!")
}
