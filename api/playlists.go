package api

import (
	"net/http"

	db "github.com/WeslleyRibeiro-1999/deejay-playlist/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPlaylistRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) createPlaylist(ctx *gin.Context) {
	var req createPlaylistRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	playlist, err := server.store.CreatePlaylist(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusCreated, playlist)
}

type getPlaylistRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getPlaylist(ctx *gin.Context) {
	var req getPlaylistRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	playlist, err := server.store.GetPlaylist(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, playlist)
}

type deletePlaylistRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deletePlaylist(ctx *gin.Context) {
	var req deletePlaylistRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeletePlaylist(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

func (server *Server) getPlaylists(ctx *gin.Context) {
	playlists, err := server.store.GetPlaylists(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, playlists)

}

type updatePlaylistRequest struct {
	ID   int32  `json:"id" binding:"required"`
	Name string `json:"name"`
}

func (server *Server) updatePlaylist(ctx *gin.Context) {
	var req updatePlaylistRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdatePlaylistParams{
		ID:   req.ID,
		Name: req.Name,
	}

	playlist, err := server.store.UpdatePlaylist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, playlist)
}
