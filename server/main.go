package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

func main() {
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.Render(http.StatusOK, render.String{Format: "Pong"})
	})

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
