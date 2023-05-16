package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Form struct {
	Colors []string `form:"colors"`
}

func formHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func formOutput(ctx *gin.Context) {
	var form Form
	ctx.ShouldBind(&form)
	ctx.JSON(http.StatusOK, gin.H{
		"colors": form.Colors,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("index.html")
	r.GET("/form", formHandler)
	r.POST("/output", formOutput)
	r.Run()
}
