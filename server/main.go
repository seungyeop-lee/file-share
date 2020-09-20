package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"os"
	"path/filepath"
)

const shareDir = "/app/file-share/share_dir"

func main() {
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.Render(http.StatusOK, render.String{Format: "Pong"})
	})

	e.GET("/list", func(c *gin.Context) {
		var files []string
		if err := filepath.Walk(
			shareDir,
			func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					relPath, _ := filepath.Rel(shareDir, path)
					files = append(files, relPath)
				}
				return nil
			},
		); err != nil {
			c.Render(http.StatusInternalServerError, render.String{Format: "can't find share dir"})
			return
		}
		c.JSON(http.StatusOK, files)
	})

	e.GET("/download/:path", func(c *gin.Context) {
		path := c.Param("path")
		fileName := filepath.Base(path)
		c.File(filepath.Join(shareDir, path))
		c.Header("Content-Disposition", fmt.Sprintf(`Content-Disposition; filename="%s"`, fileName))
	})

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
