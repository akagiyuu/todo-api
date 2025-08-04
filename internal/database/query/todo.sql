-- name: CreateTodo :one
INSERT INTO todos(account_id, title, content, priority)
VALUES ($1, $2, $3, $4)
RETURNING id;
