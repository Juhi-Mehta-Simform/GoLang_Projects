package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "0")
		name := ctx.Query("name")
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "abcd")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"id":      id,
			"name":    name,
			"message": message,
			"nick":    nick,
		})
	})
	r.Run()
}
