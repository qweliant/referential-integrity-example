package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"referentialintegrity.sqlc.dev/app/referentialintegrity"
)

func run() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close(ctx)
	queries := referentialintegrity.New(conn)
	_, err = queries.CreateApplication(ctx, referentialintegrity.CreateApplicationParams{
		Sid:      123,
		Cname:    "Clemson University",
		Major:    pgtype.Text{String: "Computer Science", Valid: true},
		Decision: pgtype.Text{String: "Accepted", Valid: true},
	})
	// this shoiuld fail because the college does not exist nor the student
	if err != nil {
		log.Println(err)
	}
	// create the student and college first
	_, err = queries.CreateStudent(ctx, referentialintegrity.CreateStudentParams{
		Sname:  pgtype.Text{String: "Test Student", Valid: true},
		Gpa:    pgtype.Float4{Float32: float32(3.5), Valid: true},
		Sizehs: pgtype.Int4{Int32: int32(1000), Valid: true},
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	run()
}
