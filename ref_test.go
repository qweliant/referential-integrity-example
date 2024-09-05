package main

import (
	"context"
	"log"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"referentialintegrity.sqlc.dev/app/referentialintegrity"
)

func TestCreateApplicationInvalidForeignKey(t *testing.T) {
	// Establish connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	require.NoError(t, err)
	defer conn.Close(ctx)

	// Initialize queries
	queries := referentialintegrity.New(conn)

	// Attempt to insert with non-existing foreign key
	_, err = queries.CreateApplication(ctx, referentialintegrity.CreateApplicationParams{
		Sid:      123,                     // Non-existent sID
		Cname:    "NonExistentUniversity", // Non-existent cName
		Major:    pgtype.Text{String: "Computer Science", Valid: true},
		Decision: pgtype.Text{String: "Accepted", Valid: true},
	})
	log.Println(err)

	// Check if error is returned as expected
	require.Error(t, err, "expected foreign key constraint violation error")
}
