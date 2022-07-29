-- name: CreateQueue :one
INSERT INTO queue(
    id_playlist,
    title_song,
    position
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetQueues :many
SELECT *
FROM queue 
INNER JOIN song 
ON title_song = title
WHERE queue.id_playlist = $1;

-- name: DeleteQueue :exec
DELETE FROM queue
WHERE id_playlist = $1 and position = $2;

-- name: ClearQueue :exec
DELETE FROM queue
WHERE id_playlist =$1;

-- name: UpdateQueue :one
UPDATE queue
SET position = $3
WHERE id_playlist = $1 AND position = $2 RETURNING *;
