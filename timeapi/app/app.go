package app

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()
	router.GET("/api/time", getCurrentTime)
	router.Run(":8090")
}
