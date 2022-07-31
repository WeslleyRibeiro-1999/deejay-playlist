package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createQueueRandom(t *testing.T) Queue {
	createdPlaylist := createPlaylistRandom(t)
	createdSong := createSongRandom(t)
	arg := CreateQueueParams{
		IDPlaylist: createdPlaylist.ID,
		TitleSong:  createdSong.Title,
		Position:   21,
	}

	queue, err := testQueries.CreateQueue(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, queue)

	require.Equal(t, queue.IDPlaylist, arg.IDPlaylist)
	require.Equal(t, queue.Position, arg.Position)
	require.Equal(t, queue.TitleSong, arg.TitleSong)

	return queue
}

func TestCreateQueue(t *testing.T) {
	createQueueRandom(t)
}
