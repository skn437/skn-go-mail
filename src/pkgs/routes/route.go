package routes

import (
	"skn-go-mail/src/pkgs/routes/mails"
	"skn-go-mail/src/pkgs/utils"

	"github.com/gin-gonic/gin"
)

func GetRoutes() {
	var router *gin.Engine = utils.GetGinRouter()

	var mailGroup string = utils.ViperEnvVariable("ROUTES.MAIL.COMMON")

	//* Routes ++
	var basicMailUrl string = utils.ViperEnvVariable("ROUTES.MAIL.BASIC")
	mails.MailRouteBasic(router, mailGroup, basicMailUrl)

	var httpMailUrl string = utils.ViperEnvVariable("ROUTES.MAIL.HTML")
	mails.HttpMailRoute(router, mailGroup, httpMailUrl)
	//* Routes --

	var port string = utils.ViperEnvVariable("PORT")
	router.Run(port)
}
