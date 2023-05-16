package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, numHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, numHandlers)
	}
	r.POST("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "foo")
	})
	r.GET("/bar", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "bar")
	})
	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	r.Run()

}
