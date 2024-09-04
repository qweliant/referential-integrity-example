// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package referentialintegrity

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createApplication = `-- name: CreateApplication :one
INSERT INTO applications (
  sID, cName, major, decision
) VALUES (
  $1, $2, $3, $4
)
RETURNING sid, cname, major, decision
`

type CreateApplicationParams struct {
	Sid      int32
	Cname    string
	Major    pgtype.Text
	Decision pgtype.Text
}

func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, createApplication,
		arg.Sid,
		arg.Cname,
		arg.Major,
		arg.Decision,
	)
	var i Application
	err := row.Scan(
		&i.Sid,
		&i.Cname,
		&i.Major,
		&i.Decision,
	)
	return i, err
}

const createCollege = `-- name: CreateCollege :one
INSERT INTO colleges (
  cName, state, enrollment
) VALUES (
  $1, $2, $3
)
RETURNING cname, state, enrollment
`

type CreateCollegeParams struct {
	Cname      string
	State      pgtype.Text
	Enrollment pgtype.Int4
}

func (q *Queries) CreateCollege(ctx context.Context, arg CreateCollegeParams) (College, error) {
	row := q.db.QueryRow(ctx, createCollege, arg.Cname, arg.State, arg.Enrollment)
	var i College
	err := row.Scan(&i.Cname, &i.State, &i.Enrollment)
	return i, err
}

const createStudent = `-- name: CreateStudent :one
INSERT INTO students (
  sID, sName, GPA, sizeHS
) VALUES (
  $1, $2, $3, $4
)
RETURNING sid, sname, gpa, sizehs
`

type CreateStudentParams struct {
	Sid    int32
	Sname  pgtype.Text
	Gpa    pgtype.Float4
	Sizehs pgtype.Int4
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRow(ctx, createStudent,
		arg.Sid,
		arg.Sname,
		arg.Gpa,
		arg.Sizehs,
	)
	var i Student
	err := row.Scan(
		&i.Sid,
		&i.Sname,
		&i.Gpa,
		&i.Sizehs,
	)
	return i, err
}
