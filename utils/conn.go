package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func ConnectToPostgres() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	u := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), u)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Postgres: ", u)

	return conn, nil
}
