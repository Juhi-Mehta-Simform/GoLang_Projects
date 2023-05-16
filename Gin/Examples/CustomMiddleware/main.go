package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Set("example", "12345")
		ctx.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := ctx.Writer.Status()
		log.Print(status)
	}
}

func main() {
	r := gin.New()
	r.Use(logger())
	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		log.Println(example)
	})
	r.Run()
}
