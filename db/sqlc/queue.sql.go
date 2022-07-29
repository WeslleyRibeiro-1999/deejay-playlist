// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: queue.sql

package db

import (
	"context"
)

const clearQueue = `-- name: ClearQueue :exec
DELETE FROM queue
WHERE id_playlist =$1
`

func (q *Queries) ClearQueue(ctx context.Context, idPlaylist int32) error {
	_, err := q.db.ExecContext(ctx, clearQueue, idPlaylist)
	return err
}

const createQueue = `-- name: CreateQueue :one
INSERT INTO queue(
    id_playlist,
    title_song,
    position
) VALUES (
    $1, $2, $3
) RETURNING id_playlist, title_song, position
`

type CreateQueueParams struct {
	IDPlaylist int32  `json:"id_playlist"`
	TitleSong  string `json:"title_song"`
	Position   int32  `json:"position"`
}

func (q *Queries) CreateQueue(ctx context.Context, arg CreateQueueParams) (Queue, error) {
	row := q.db.QueryRowContext(ctx, createQueue, arg.IDPlaylist, arg.TitleSong, arg.Position)
	var i Queue
	err := row.Scan(&i.IDPlaylist, &i.TitleSong, &i.Position)
	return i, err
}

const deleteQueue = `-- name: DeleteQueue :exec
DELETE FROM queue
WHERE id_playlist = $1 and position = $2
`

type DeleteQueueParams struct {
	IDPlaylist int32 `json:"id_playlist"`
	Position   int32 `json:"position"`
}

func (q *Queries) DeleteQueue(ctx context.Context, arg DeleteQueueParams) error {
	_, err := q.db.ExecContext(ctx, deleteQueue, arg.IDPlaylist, arg.Position)
	return err
}

const getQueues = `-- name: GetQueues :many
SELECT id_playlist, title_song, position, title, genre, artist, duration
FROM queue 
INNER JOIN song 
ON title_song = title
WHERE queue.id_playlist = $1
`

type GetQueuesRow struct {
	IDPlaylist int32  `json:"id_playlist"`
	TitleSong  string `json:"title_song"`
	Position   int32  `json:"position"`
	Title      string `json:"title"`
	Genre      string `json:"genre"`
	Artist     string `json:"artist"`
	Duration   int32  `json:"duration"`
}

func (q *Queries) GetQueues(ctx context.Context, idPlaylist int32) ([]GetQueuesRow, error) {
	rows, err := q.db.QueryContext(ctx, getQueues, idPlaylist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetQueuesRow{}
	for rows.Next() {
		var i GetQueuesRow
		if err := rows.Scan(
			&i.IDPlaylist,
			&i.TitleSong,
			&i.Position,
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

const updateQueue = `-- name: UpdateQueue :one
UPDATE queue
SET position = $3
WHERE id_playlist = $1 AND position = $2 RETURNING id_playlist, title_song, position
`

type UpdateQueueParams struct {
	IDPlaylist int32 `json:"id_playlist"`
	Position   int32 `json:"position"`
	Position_2 int32 `json:"position_2"`
}

func (q *Queries) UpdateQueue(ctx context.Context, arg UpdateQueueParams) (Queue, error) {
	row := q.db.QueryRowContext(ctx, updateQueue, arg.IDPlaylist, arg.Position, arg.Position_2)
	var i Queue
	err := row.Scan(&i.IDPlaylist, &i.TitleSong, &i.Position)
	return i, err
}