-- name: CreatePlaylist :one
INSERT INTO playlist (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetPlaylist :one
SELECT * FROM playlist
WHERE id = $1 LIMIT 1;

-- name: GetPlaylists :many
SELECT * FROM playlist;

-- name: UpdatePlaylist :one
UPDATE playlist
SET name = $2
WHERE id = $1 RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlist
WHERE id = $1;
