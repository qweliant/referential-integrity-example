-- name: CreateApplication :one
INSERT INTO applications (
  sID, cName, major, decision
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: CreateStudent :one
INSERT INTO students (
  sID, sName, GPA, sizeHS
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: CreateCollege :one
INSERT INTO colleges (
  cName, state, enrollment
) VALUES (
  $1, $2, $3
)
RETURNING *;
