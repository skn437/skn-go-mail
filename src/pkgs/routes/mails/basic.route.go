package mails

import (
	"io"

	"github.com/gin-gonic/gin"
)

func MailRouteBasic(router *gin.Engine, url string) {
	router.POST(url, func(ctx *gin.Context) {
		var body io.ReadCloser = ctx.Request.Body
		value, err := io.ReadAll(body)

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, gin.H{
			"message": string(value),
		})
	})

	router.GET(url, func(ctx *gin.Context) {
		ctx.String(200, "Email Sent Successfully!")
	})
}
