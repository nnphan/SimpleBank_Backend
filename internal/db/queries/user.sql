-- name: CreateUser :one
INSERT INTO users (
  id,
  full_name,
  email,
  password_hash
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE full_name = $1 LIMIT 1;


-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
