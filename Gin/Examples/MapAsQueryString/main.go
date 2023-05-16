package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/post", func(ctx *gin.Context) {
		id := ctx.QueryMap("id")
		name := ctx.PostFormMap("name")
		fmt.Printf("id: %v, name: %v", id, name)

	})
	r.Run(":8080")
}
