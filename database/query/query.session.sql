-- name: CreateSession :one
INSERT INTO sessions (
  id,
  uid,
  username,
  refresh_token,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;