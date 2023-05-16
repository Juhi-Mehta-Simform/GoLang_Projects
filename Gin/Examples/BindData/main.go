package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	A      StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	A      *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	X struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetB(ctx *gin.Context) {
	var b StructB
	ctx.Bind(&b)
	ctx.JSON(http.StatusOK, gin.H{
		"a": b.A,
		"b": b.FieldB,
	})
}

func GetC(ctx *gin.Context) {
	var c StructC
	ctx.Bind(&c)
	ctx.JSON(http.StatusOK, gin.H{
		"a": c.A,
		"c": c.FieldC,
	})
}

func GetD(ctx *gin.Context) {
	var d StructD
	ctx.Bind(&d)
	ctx.JSON(http.StatusOK, gin.H{
		"a": d.X,
		"c": d.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/bindB", GetB)
	r.GET("/bindC", GetC)
	r.GET("/bindD", GetD)
	r.Run()
}
