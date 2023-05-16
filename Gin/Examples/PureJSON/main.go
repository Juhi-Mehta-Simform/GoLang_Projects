package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello</b>",
		})
	})
	r.GET("/pure", func(ctx *gin.Context) {
		ctx.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello</b>",
		})
	})
	r.Run()
}
