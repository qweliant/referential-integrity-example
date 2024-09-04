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

-- name: UpdateApplication :one
UPDATE applications
SET major = $2, decision = $3
WHERE sID = $1 AND cName = $4
RETURNING *;

-- name: UpdateStudent :one
UPDATE students
SET sName = $2, GPA = $3, sizeHS = $4
WHERE sID = $1
RETURNING *;

-- name: UpdateCollege :one
UPDATE colleges
SET state = $2, enrollment = $3
WHERE cName = $1
RETURNING *;

-- name: DeleteApplication :one
DELETE FROM applications
WHERE sID = $1 AND cName = $2
RETURNING *;

-- name: DeleteStudent :one
DELETE FROM students
WHERE sID = $1
RETURNING *;

-- name: DeleteCollege :one
DELETE FROM colleges
WHERE cName = $1
RETURNING *;

