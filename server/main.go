package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.Render(http.StatusOK, render.String{Format: "Pong"})
	})

	e.GET("/list", func(c *gin.Context) {
		var files []string
		if err := filepath.Walk(
			"/app/file-share/share_dir",
			func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					files = append(files, path)
				}
				return nil
			},
		); err != nil {
			c.Render(http.StatusInternalServerError, render.String{Format: "can't find share dir"})
			return
		}
		c.JSON(http.StatusOK, files)
	})

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
