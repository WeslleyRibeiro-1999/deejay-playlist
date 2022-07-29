-- name: CreateSong :one
INSERT INTO song (
    title,
    genre,
    artist,
    duration
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSongs :many
SELECT * FROM song;

-- name: DeleteSong :exec
DELETE FROM song
WHERE title = $1;
