package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)


func Run() error {
	mux := gin.Default()

	v1 := mux.Group("/api/v1/")
	{
		v1.GET("/test" , func(context *gin.Context) {
			context.JSON(200 , gin.H{"message":"ok"})
		})
	}

	mux.Use(static.Serve("/", static.LocalFile("../front/dist", false)))

	return mux.Run(":4000")
}