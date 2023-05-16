package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/foo", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "foo")
		})
		v1.POST("/bar", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "bar")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/foo", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "foo")
		})
		v2.POST("/bar", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "bar")
		})
	}
	r.Run()
}
