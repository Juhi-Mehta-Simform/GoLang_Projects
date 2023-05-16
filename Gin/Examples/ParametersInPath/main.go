package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/param/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := "Hello " + name + " , " + action
		ctx.JSON(http.StatusOK, message)
	})
	r.Run()
}
