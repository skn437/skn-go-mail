package routes

import (
	"skn-go-mail/src/pkgs/routes/mails"
	"skn-go-mail/src/pkgs/utils"

	"github.com/gin-gonic/gin"
)

func GetRoutes() {
	var router *gin.Engine = utils.GetGinRouter()

	var mailBasic string = utils.ViperEnvVariable("ROUTES.MAIL.BASIC")
	mails.MailRouteBasic(router, mailBasic)

	var port string = utils.ViperEnvVariable("PORT")
	router.Run(port)
}
