package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name     string `form:"name",binding:"required"`
	Password int    `form:"password",binding:"required"`
}

func main() {
	var form Login
	r := gin.Default()
	r.POST("/login", func(ctx *gin.Context) {
		if ctx.ShouldBind(&form) == nil {
			if form.Name == "juhi" && form.Password == 1234 {
				ctx.JSON(http.StatusOK, gin.H{
					"status": "You are logged in.",
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status": "Unauthorized",
				})
			}

		}
	})
	r.Run()
}
