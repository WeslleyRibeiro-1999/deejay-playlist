package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createPlaylistRandom(t *testing.T) Playlist {
	playlist, err := testQueries.CreatePlaylist(context.Background(), "pop")

	require.NoError(t, err)
	require.NotEmpty(t, playlist)

	require.NotEmpty(t, playlist.ID)
	require.NotEmpty(t, playlist.Name)
	require.NotEmpty(t, playlist.CreatedAt)
	require.Equal(t, "pop", playlist.Name)

	return playlist
}

func TestCreatedPlaylist(t *testing.T) {
	createPlaylistRandom(t)
}

func TestDeletedPlaylist(t *testing.T) {
	createdPlaylist := createPlaylistRandom(t)

	err := testQueries.DeletePlaylist(context.Background(), createdPlaylist.ID)

	require.NoError(t, err)
}

func TestGetPlaylist(t *testing.T) {
	createPlaylist := createPlaylistRandom(t)

	playlist, err := testQueries.GetPlaylist(context.Background(), createPlaylist.ID)

	require.NoError(t, err)
	require.NotEmpty(t, playlist)

	require.Equal(t, createPlaylist.ID, playlist.ID)
	require.Equal(t, createPlaylist.Name, playlist.Name)
	require.Equal(t, createPlaylist.CreatedAt, playlist.CreatedAt)
}

func TestGetPlaylists(t *testing.T) {
	createPlaylist := createPlaylistRandom(t)

	finded, err := testQueries.GetPlaylists(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, createPlaylist)

	for _, playlists := range finded {
		require.NotEmpty(t, playlists.ID)
		require.NotEmpty(t, playlists.Name)
		require.NotEmpty(t, playlists.CreatedAt)
	}
}

func TestUpdatePlaylist(t *testing.T) {
	createPlaylist := createPlaylistRandom(t)

	arg := UpdatePlaylistParams{
		ID:   createPlaylist.ID,
		Name: "rock",
	}

	playlistUpdated, err := testQueries.UpdatePlaylist(context.Background(), arg)

	require.NotEmpty(t, createPlaylist)
	require.NotEmpty(t, playlistUpdated)
	require.NoError(t, err)

	require.Equal(t, playlistUpdated.ID, createPlaylist.ID)
	require.Equal(t, arg.Name, playlistUpdated.Name)
	require.Equal(t, createPlaylist.CreatedAt, playlistUpdated.CreatedAt)
}
