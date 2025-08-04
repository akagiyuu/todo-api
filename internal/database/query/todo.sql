-- name: CreateTodo :one
INSERT INTO todos(account_id, title, content, priority)
VALUES (@account_id, @title, @content, @priority)
RETURNING id;

-- name: GetTodo :one
SELECT title, content, priority, is_done, created_at
FROM todos
WHERE id = @id AND account_id = @account_id;
