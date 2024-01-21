package middlewares

import "github.com/gin-gonic/gin"

// * If an argument is needed to pass then middlewares should be written like this one
func CsrfMiddleware(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("X-CSRF-TOKEN") != token {
			ctx.AbortWithStatusJSON(500, gin.H{
				"message": "CSRF Token Is Invalid Or Not Present",
			})

			return
		}

		ctx.Next()
	}
}

// * This middleware adds a header to the response
func AddHeader(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("skn-key", key)

		ctx.Next()
	}
}
