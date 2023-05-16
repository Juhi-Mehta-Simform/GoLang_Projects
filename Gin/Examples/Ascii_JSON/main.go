package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/JSON", func(ctx *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		ctx.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8080")
}

//AsciiJSON converts the non-ascii characters to utf-16 encoding.
