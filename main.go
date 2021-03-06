package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {

	connStr := "postgresql://docker:docker@127.0.0.1:5432/docker"

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	type Person struct {
		Name string
		Age  int64
	}

	var p Person

	err = conn.QueryRow(context.Background(), "select name, age from Person").Scan(&p.Name, &p.Age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(p)
}
