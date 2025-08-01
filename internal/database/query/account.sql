-- name: CreateAccount :one
INSERT INTO accounts(email, password)
VALUES ($1, $2)
RETURNING id;
