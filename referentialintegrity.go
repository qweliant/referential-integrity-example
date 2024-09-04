package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"

	"referentialintegrity.sqlc.dev/app/referentialintegrity"
)



func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := referentialintegrity.New(conn)

	
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}