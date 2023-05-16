package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/async", func(ctx *gin.Context) {
		copy := ctx.Copy() // while using goruitnes we SHOULD NOT use original context, so make a copy of that
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done!!!" + copy.Request.URL.Path)
		}()
	})

	r.GET("/sync", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done!!!", ctx.Request.URL.Path)
	})
	r.Run()
}
