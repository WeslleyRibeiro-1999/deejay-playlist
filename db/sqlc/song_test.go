package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createSongRandom(t *testing.T) Song {
	arg := CreateSongParams{
		Title:    "Any",
		Genre:    "pop",
		Artist:   "any_name",
		Duration: 12345,
	}

	song, err := testQueries.CreateSong(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, song)
	require.Equal(t, song.Title, arg.Title)
	require.Equal(t, song.Genre, arg.Genre)
	require.Equal(t, song.Artist, arg.Artist)
	require.Equal(t, song.Duration, arg.Duration)

	return song
}

func TestCreatedSong(t *testing.T) {
	createSongRandom(t)
}

func TestDeleteSong(t *testing.T) {
	createSongRandom(t)

	err := testQueries.DeleteSong(context.Background(), "Any")

	require.NoError(t, err)
}

func TestGetSongs(t *testing.T) {
	createdSong := createSongRandom(t)

	song, err := testQueries.GetSongs(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, createdSong)

	for _, songs := range song {
		require.NotEmpty(t, songs.Title)
		require.NotEmpty(t, songs.Genre)
		require.NotEmpty(t, songs.Artist)
		require.NotEmpty(t, songs.Duration)
	}
}
