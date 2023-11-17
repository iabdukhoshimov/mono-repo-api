-- name: ListUser :many
SELECT * FROM "user";

-- name: CountUser :one
SELECT count(*) FROM "user";

-- name: CreateUser :exec
INSERT INTO "user" (name, email) VALUES ($1, $2);

-- name: GetUserById :one
SELECT * FROM "user"
WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM "user"
WHERE name = $1;