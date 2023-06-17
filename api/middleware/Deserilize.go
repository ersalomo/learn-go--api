package middleware

import (
	"github.com/gin-gonic/gin"
)

func DeserializeMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//err := jwt.TokenValid(context)
		//if err != nil {
		//	context.String(http.StatusUnauthorized, "Unauthorized")
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}
