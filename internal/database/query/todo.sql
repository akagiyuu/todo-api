-- name: CreateTodo :one
INSERT INTO todos(account_id, title, content, priority)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetTodo :one
SELECT id, title, content, priority, is_done, created_at
FROM todos
WHERE id = $1 AND account_id = $2;
