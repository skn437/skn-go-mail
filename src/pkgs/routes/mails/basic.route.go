package routes_mails

import (
	"github.com/gin-gonic/gin"
)

func MailRouteBasic(router *gin.Engine, url string) {
	router.POST(url, func(ctx *gin.Context) {
		ctx.String(200, "Email Sent Successfully!")
	})
}
