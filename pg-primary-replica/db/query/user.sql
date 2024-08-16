-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;

-- name: GetUserByID :one
SELECT id, name, created_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, created_at
FROM users
ORDER BY created_at DESC;

-- name: Updatename :one
UPDATE users
SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;
