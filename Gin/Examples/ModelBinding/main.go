package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name     string `json:"name"`
	Password int    `json:"password"`
}

func main() {
	r := gin.Default()
	var login Login
	r.POST("/login", func(ctx *gin.Context) {
		err := ctx.BindJSON(&login)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "You are successfully logged in.",
			})
			//ctx.String(http.StatusOK, "You are successfully logged in.")
		}
	})
	r.Run()
}
