package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("Cookies")
		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("Cookies", "test", 60, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s", cookie)
	})
	r.Run()
}
