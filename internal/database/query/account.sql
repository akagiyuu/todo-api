-- name: CreateAccount :one
INSERT INTO accounts(email, password)
VALUES ($1, $2)
RETURNING id;

-- name: GetAccountByEmail :one
SELECT id, password
FROM accounts
WHERE email = $1;
