package main

import (
	"nix-service-go/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/spacex/launches", handler.LaunchesGet())
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}