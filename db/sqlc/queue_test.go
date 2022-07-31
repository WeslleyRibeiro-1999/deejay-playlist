package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createQueueRandom(t *testing.T) Queue {
	createdPlaylist := createPlaylistRandom(t)
	arg := CreateQueueParams{
		IDPlaylist: createdPlaylist.ID,
		TitleSong:  "Any",
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

func TestDeleteQueue(t *testing.T) {
	createdQueue := createQueueRandom(t)
	arg := DeleteQueueParams{
		IDPlaylist: createdQueue.IDPlaylist,
		Position:   createdQueue.Position,
	}

	err := testQueries.DeleteQueue(context.Background(), arg)

	require.NoError(t, err)
}

func TestClearQueue(t *testing.T) {
	createdQueue := createQueueRandom(t)

	err := testQueries.ClearQueue(context.Background(), createdQueue.IDPlaylist)

	require.NoError(t, err)
}

func TestGetQueues(t *testing.T) {
	createdQueue := createQueueRandom(t)

	queue, err := testQueries.GetQueues(context.Background(), createdQueue.IDPlaylist)

	require.NoError(t, err)
	require.NotEmpty(t, queue)

	for _, queues := range queue {
		require.NotEmpty(t, queues.Artist)
		require.NotEmpty(t, queues.IDPlaylist)
		require.NotEmpty(t, queues.Genre)
		require.NotEmpty(t, queues.Duration)
		require.NotEmpty(t, queues.Position)
		require.NotEmpty(t, queues.TitleSong)
	}
}

func TestUpdateQueue(t *testing.T) {
	createdQueue := createQueueRandom(t)

	arg := UpdateQueueParams{
		IDPlaylist: createdQueue.IDPlaylist,
		Position:   createdQueue.Position,
		Position_2: 33,
	}

	newPosition, err := testQueries.UpdateQueue(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, createdQueue)
	require.NotEmpty(t, newPosition)

	require.Equal(t, newPosition.IDPlaylist, createdQueue.IDPlaylist)
	require.Equal(t, arg.Position_2, newPosition.Position)

}
