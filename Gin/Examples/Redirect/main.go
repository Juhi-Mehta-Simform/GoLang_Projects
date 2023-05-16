package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test1", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "https://www.google.com/")
	})
	r.GET("/foo", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/bar")
		//ctx.Request.URL.Path = "/bar"
		//r.HandleContext(ctx)
	})
	r.GET("/bar", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.Run()
}
