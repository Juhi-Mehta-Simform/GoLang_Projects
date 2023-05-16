package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(ctx *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})
	r.Run()
}
