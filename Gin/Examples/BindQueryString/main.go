package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(ctx *gin.Context) {
	var person Person
	if ctx.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		ctx.String(200, "Success")
	}

}

func main() {
	r := gin.Default()
	r.Any("/query", startPage)
	r.Run(":8080")
}
