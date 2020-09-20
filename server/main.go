package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

const (
	shareDir      = "/app/file-share/share_dir"
	basicAuthUser = "BASIC_AUTH_USER"
	basicAuthPass = "BASIC_AUTH_PASS"
)

func main() {
	e := gin.Default()
	e.Use(gin.BasicAuth(gin.Accounts{
		os.Getenv(basicAuthUser): os.Getenv(basicAuthPass),
	}))
	e.LoadHTMLGlob("templates/*")

	e.GET("/ping", func(c *gin.Context) {
		c.Render(http.StatusOK, render.String{Format: "Pong"})
	})

	e.GET("/list", func(c *gin.Context) {
		var files []string
		if err := filepath.Walk(
			shareDir,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if ext := filepath.Ext(path); ext == ".srt" || ext == ".smi" {
					return nil
				}
				relPath, _ := filepath.Rel(shareDir, path)
				files = append(files, relPath)
				return nil
			},
		); err != nil {
			c.Render(http.StatusInternalServerError, render.String{Format: "can't find share dir"})
			return
		}
		c.HTML(http.StatusOK, "list.gohtml", gin.H{
			"files": files,
		})
	})

	e.GET("/download", func(c *gin.Context) {
		path := c.Query("path")
		fileName := filepath.Base(path)
		c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
		c.File(filepath.Join(shareDir, path))
	})

	e.GET("/streaming", func(c *gin.Context) {
		path := c.Query("path")
		c.File(filepath.Join(shareDir, path))
	})

	e.GET("/play", func(c *gin.Context) {
		path := c.Query("path")
		c.HTML(http.StatusOK, "play.gohtml", gin.H{
			"path": path,
		})
	})

	e.GET("/sub", func(c *gin.Context) {
		path := c.Query("path")
		ext := filepath.Ext(path)

		smiPath := filepath.Join(shareDir, strings.Replace(path, ext, ".smi", -1))
		if sub, err := astisub.OpenFile(smiPath); err == nil {
			err = sub.WriteToWebVTT(c.Writer)
			c.Status(http.StatusOK)
			return
		}

		srtPath := filepath.Join(shareDir, strings.Replace(path, ext, ".srt", -1))
		if sub, err := astisub.OpenFile(srtPath); err == nil {
			err = sub.WriteToWebVTT(c.Writer)
			c.Status(http.StatusOK)
			return
		}

		c.Status(http.StatusInternalServerError)
	})

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
