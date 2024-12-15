package router

import (
	"io"
	cors "itish2003/image-primitive/middleware"
	"itish2003/image-primitive/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	file, err := os.Create("logFile.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router.Use(gin.Logger())
	router.Use(cors.CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Base endpoint",
			})
		})

	route := router.Group("/v1")
	{
		route.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Working endpoint",
			})
		})
		route.POST("/upload", service.UploadImage)
	}
	return router
}
