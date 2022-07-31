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

	//AQUI FICA AS ROTAS

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api started with error:": err.Error()}
}
