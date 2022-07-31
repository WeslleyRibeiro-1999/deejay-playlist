package api

import (
	db "github.com/WeslleyRibeiro-1999/deejay-playlist/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func InstanceServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//AQUI FICA AS ROTAS PLAYLISTS
	router.POST("/playlist/:name", server.createPlaylist)
	router.GET("/playlist/:id", server.getPlaylist)
	router.GET("/playlists", server.getPlaylists)
	router.DELETE("/playlist/:id", server.deletePlaylist)
	router.PUT("/playlist", server.updatePlaylist)

	//AQUI FICA AS ROTAS SONGS
	router.POST("/song", server.createSong)
	router.GET("/songs", server.getSongs)
	router.DELETE("/song/:title", server.deleteSong)

	//AQUI FICA AS ROTAS QUEUES
	router.POST("/queue", server.createQueue)
	router.GET("/queue/:id", server.getQueue)
	router.DELETE("/queue", server.deleteQueue)
	router.PUT("/queue", server.updateQueue)
	router.DELETE("queue/:id", server.clearQueue)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api started with error:": err.Error()}
}
