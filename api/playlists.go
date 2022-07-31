package api

import (
	"net/http"

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
