package server

import (
	"github.com/VusalShahbazov/chat-app/app/pkg/models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Handler *gin.Engine
	StaticFilePath string
}

func Run() error {
	server := &Server{
		Handler: gin.Default(),
		StaticFilePath: "../front/dist",
	}

	//api routes
	server.DefineRoutes()

	//vue bundle
	server.CreateStaticFileServer()

	err := models.ConnectToDb()
	if err != nil  {
		return err
	}

	return server.Handler.Run(":4000")
}


func (s *Server) CreateStaticFileServer()  {
	s.Handler.Use(static.Serve("/", static.LocalFile(s.StaticFilePath, false)))
}

