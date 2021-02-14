package server

import (
	"github.com/gin-gonic/gin"
)


func (s *Server) DefineRoutes()  {

	v1 := s.Handler.Group("/api/v1/")
	{
		v1.GET("/test" , func(context *gin.Context) {
			context.JSON(200 , gin.H{"message":"ok"})
		})
	}
}

