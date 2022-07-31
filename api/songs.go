package api

import (
	"net/http"

	db "github.com/WeslleyRibeiro-1999/deejay-playlist/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSongRequest struct {
	Title    string `json:"title" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
	Artist   string `json:"artist" binding:"required"`
	Duration int32  `json:"duration" binding:"required"`
}

func (server *Server) createSong(ctx *gin.Context) {
	var req createSongRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateSongParams{
		Title:    req.Title,
		Genre:    req.Genre,
		Artist:   req.Artist,
		Duration: req.Duration,
	}

	song, err := server.store.CreateSong(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusCreated, song)

}

type deleteSongRequest struct {
	Title string `uri:"title" binding:"required"`
}

func (server *Server) deleteSong(ctx *gin.Context) {
	var req deleteSongRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteSong(ctx, req.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

func (server *Server) getSongs(ctx *gin.Context) {
	songs, err := server.store.GetSongs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, songs)

}
