package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.SecureJsonPrefix(")]}',\n")
	//r.SecureJsonPrefix("juhi,")
	r.GET("/secure", func(ctx *gin.Context) {
		name := []string{"foo", "bar", "json"}
		ctx.SecureJSON(http.StatusOK, name)
	})
	r.Run()
}
