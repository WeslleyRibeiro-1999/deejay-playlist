package api

import (
	"net/http"

	db "github.com/WeslleyRibeiro-1999/deejay-playlist/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createQueueRequest struct {
	IDPlaylist int32  `json:"id_playlist" binding:"required"`
	TitleSong  string `json:"title_song" binding:"required"`
	Position   int32  `json:"position" binding:"required"`
}

func (server *Server) createQueue(ctx *gin.Context) {
	var req createQueueRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateQueueParams{
		IDPlaylist: req.IDPlaylist,
		TitleSong:  req.TitleSong,
		Position:   req.Position,
	}

	queue, err := server.store.CreateQueue(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusCreated, queue)
}

type clearQueueRequest struct {
	ID_Playlist int32 `uri:"id_playlist" binding:"required"`
}

func (server *Server) clearQueue(ctx *gin.Context) {
	var req clearQueueRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.ClearQueue(ctx, req.ID_Playlist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

type deleteQueueRequest struct {
	IDPlaylist int32 `json:"id_playlist" binding:"required"`
	Position   int32 `json:"position" binding:"required"`
}

func (server *Server) deleteQueue(ctx *gin.Context) {
	var req deleteQueueRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.DeleteQueueParams{
		IDPlaylist: req.IDPlaylist,
		Position:   req.Position,
	}

	err = server.store.DeleteQueue(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

type getQueueRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getQueue(ctx *gin.Context) {
	var req getQueueRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	queues, err := server.store.GetQueues(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, queues)
}

type updateQueueRequest struct {
	IDPlaylist int32 `json:"id_playlist" binding:"required"`
	Position   int32 `json:"position" binding:"required"`
	Position_2 int32 `json:"position_2" binding:"required"`
}

func (server *Server) updateQueue(ctx *gin.Context) {
	var req updateQueueRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateQueueParams{
		IDPlaylist: req.IDPlaylist,
		Position:   req.Position,
		Position_2: req.Position_2,
	}

	playlist, err := server.store.UpdateQueue(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, playlist)
}
