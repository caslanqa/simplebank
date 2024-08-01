-- name: CreateUser :one
-- burada query tanımlamaları yapılıyor
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) values (
    $1,$2,$3,$4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;