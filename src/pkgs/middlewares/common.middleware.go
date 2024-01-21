package middlewares

import "github.com/gin-gonic/gin"

func CommonMiddleware(ctx *gin.Context) {
	if ctx.Request.Header.Get("token") != "SKN" {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Token Not Found"},
		)

		return
	}

	ctx.Next()
}
