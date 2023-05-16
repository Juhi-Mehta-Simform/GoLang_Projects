package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   int    `uri:"id"`
	Name string `uri:"name"`
}

func bindUri(ctx *gin.Context) {
	var person Person
	err := ctx.ShouldBindUri(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"uuid": person.ID,
		})
	}
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", bindUri)
	r.Run()
}
