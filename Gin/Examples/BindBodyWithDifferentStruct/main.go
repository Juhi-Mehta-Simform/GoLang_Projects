package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type FormA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type FormB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()
	var a FormA
	var b FormB
	r.POST("/json", func(ctx *gin.Context) {
		if ctx.ShouldBindBodyWith(&a, binding.JSON) == nil {
			ctx.String(http.StatusOK, "Body should be FormA, JSON")
		} else if ctx.ShouldBindBodyWith(&b, binding.JSON) == nil {
			ctx.String(http.StatusOK, "Body should be FormB, JSON")
		} else if ctx.ShouldBindBodyWith(&a, binding.XML) == nil {
			ctx.String(http.StatusOK, "Body should be FormA, XML")
		} else {
			ctx.String(http.StatusBadRequest, "Error")
		}
	})
	r.Run()
}
