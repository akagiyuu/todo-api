-- name: CreateTodo :one
INSERT INTO todos(account_id, title, content, priority)
VALUES (@account_id, @title, @content, @priority)
RETURNING id;

-- name: GetTodo :one
SELECT title, content, priority, is_done, created_at
FROM todos
WHERE id = @id AND account_id = @account_id;

-- name: FilterTodo :many
SELECT id, title, content, priority, is_done, created_at
FROM todos
WHERE account_id = @account_id AND
    (
        @query IS NULL OR
        title LIKE '%' || @query || '%' OR
        content LIKE '%' || @query || '%'
    ) AND
    (@priority IS NULL OR priority = @priority) AND
    (@is_done IS NULL OR is_done = @is_done);

-- name: UpdateTodo :exec
UPDATE todos
SET
    title = COALESCE(@title, title),
    content = COALESCE(@content, content),
    priority = COALESCE(@priority, priority)
WHERE id = @id AND account_id = @account_id AND is_done = true;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = @id AND account_id = @account_id;
