package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed static/index.html
var index string

func main() {

	r := gin.Default()

	r.GET("/idk", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(index))
	})

	r.Run(":8080")

}
