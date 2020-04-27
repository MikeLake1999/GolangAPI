package middleware

import (
	"errors"
	"gallery/services"
	"strings"

	"github.com/gin-gonic/gin"
)

//check authentication
func RequireAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		arr := strings.Split(token, " ")

		if len(arr) != 2 {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return

		}

		if arr[0] != "Bearer" {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		token = arr[1]
		claims, err := services.VerifyToken(token)

		if err != nil {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		ctx.Set("account_id", claims.AccountID)

		ctx.Next()
	}
}
