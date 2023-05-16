package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ping")
	})
	r.Run()
}
