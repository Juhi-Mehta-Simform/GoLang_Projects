package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/Gin/service"
	"log"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		fmt.Println(tokenString)
		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {

		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
