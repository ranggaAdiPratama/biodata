-- name: GetAllUser :many

SELECT * FROM users ORDER BY id;
-- name: GetUser :one

SELECT * FROM users WHERE id = $1 LIMIT 1;
-- name: GetUserForUpdate :one

SELECT * FROM users WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetUserByUsername :one

SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreateUser :one

INSERT INTO
    users (
        username, name, email, password
    )
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one

UPDATE users
SET
    username = $2,
    name = $3,
    email = $4,
    profile_picture = $5,
    updated_at = $6
WHERE
    id = $1 RETURNING *;

-- name: UpdateUserPassword :one

UPDATE users
SET password = $2, updated_at = $3
WHERE
    id = $1 RETURNING *;