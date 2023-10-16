package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func GetConnection() *pgxpool.Pool {
	dbURL := "user=postgres password=admin host=127.0.0.1 port=5432 dbname=latihan pool_max_conns=5 pool_max_conn_lifetime=3s pool_max_conn_idle_time=3s"
	dbpool, poolErr := pgxpool.New(context.Background(), dbURL)
	if poolErr != nil {
		fmt.Println("Error saat membuat connection pool :", poolErr.Error())
		os.Exit(1)
	}

	err := dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database : %v\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Successfully connect to database")
	}

	return dbpool
}
