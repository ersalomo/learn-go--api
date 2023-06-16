package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api/jwt"
	"net/http"
)

func DeserilizeMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := jwt.TokenValid(context)
		if err != nil {
			context.String(http.StatusUnauthorized, "Unathorized")
			context.Abort()
			return
		}
		context.Next()
	}
}
