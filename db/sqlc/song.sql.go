// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: song.sql

package db

import (
	"context"
)

const createSong = `-- name: CreateSong :one
INSERT INTO song (
    title,
    genre,
    artist,
    duration
) VALUES (
    $1, $2, $3, $4
) RETURNING title, genre, artist, duration
`

type CreateSongParams struct {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Artist   string `json:"artist"`
	Duration int32  `json:"duration"`
}

func (q *Queries) CreateSong(ctx context.Context, arg CreateSongParams) (Song, error) {
	row := q.db.QueryRowContext(ctx, createSong,
		arg.Title,
		arg.Genre,
		arg.Artist,
		arg.Duration,
	)
	var i Song
	err := row.Scan(
		&i.Title,
		&i.Genre,
		&i.Artist,
		&i.Duration,
	)
	return i, err
}

const deleteSong = `-- name: DeleteSong :exec
DELETE FROM song
WHERE title = $1
`

func (q *Queries) DeleteSong(ctx context.Context, title string) error {
	_, err := q.db.ExecContext(ctx, deleteSong, title)
	return err
}

const getSongs = `-- name: GetSongs :many
SELECT title, genre, artist, duration FROM song
`

func (q *Queries) GetSongs(ctx context.Context) ([]Song, error) {
	rows, err := q.db.QueryContext(ctx, getSongs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Song{}
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.Title,
			&i.Genre,
			&i.Artist,
			&i.Duration,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
