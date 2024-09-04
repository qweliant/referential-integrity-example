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
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
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
	if err != nil {
		log.Println(err)
	}
	_, err = queries.CreateApplication(ctx, referentialintegrity.CreateApplicationParams{
		Sid:      234,
		Cname:    "Georgia Tech",
		Major:    pgtype.Text{String: "Computer Science", Valid: true},
		Decision: pgtype.Text{String: "Denied", Valid: true},
	})
	if err != nil {
		log.Println(err)
	}

	// both the prvious queries should fail because the student and college do not exist

	// create the student and college first
	_, err = queries.CreateStudent(ctx, referentialintegrity.CreateStudentParams{
		Sid:    123,
		Sname:  pgtype.Text{String: "Demetrius", Valid: true},
		Gpa:    pgtype.Float4{Float32: float32(3.5), Valid: true},
		Sizehs: pgtype.Int4{Int32: int32(1000), Valid: true},
	})
	if err != nil {
		log.Println(err)
	}
	_, err = queries.CreateStudent(ctx, referentialintegrity.CreateStudentParams{
		Sid:    234,
		Sname:  pgtype.Text{String: "Ananstasia", Valid: true},
		Gpa:    pgtype.Float4{Float32: float32(3.5), Valid: true},
		Sizehs: pgtype.Int4{Int32: int32(1000), Valid: true},
	})
	if err != nil {
		log.Println(err)
	}

	_, err = queries.CreateCollege(ctx, referentialintegrity.CreateCollegeParams{
		Cname: "Clemson University",
		State: pgtype.Text{String: "SC", Valid: true},
		Enrollment: pgtype.Int4{
			Int32: 10000,
			Valid: true,
		},
	})
	if err != nil {
		log.Println(err)
	}

	_, err = queries.CreateCollege(ctx, referentialintegrity.CreateCollegeParams{
		Cname: "Georgia Tech",
		State: pgtype.Text{String: "SC", Valid: true},
		Enrollment: pgtype.Int4{
			Int32: 10000,
			Valid: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
	_, err = queries.CreateApplication(ctx, referentialintegrity.CreateApplicationParams{
		Sid:      123,
		Cname:    "Clemson University",
		Major:    pgtype.Text{String: "Computer Science", Valid: true},
		Decision: pgtype.Text{String: "Accepted", Valid: true},
	})
	if err != nil {
		log.Println(err)
	}
	_, err = queries.CreateApplication(ctx, referentialintegrity.CreateApplicationParams{
		Sid:      234,
		Cname:    "Georgia Tech",
		Major:    pgtype.Text{String: "Computer Science", Valid: true},
		Decision: pgtype.Text{String: "Denied", Valid: true},
	})
	if err != nil {
		log.Println(err)
	}

	// update the apply table
	_, err = queries.UpdateApplication(ctx, referentialintegrity.UpdateApplicationParams{
		Sid: 345,
	})

}

func main() {
	run()
}
