package routes

import (
	routes_mails "skn-go-mail/src/pkgs/routes/mails"
	utils_functions "skn-go-mail/src/pkgs/utils/functions"

	"github.com/gin-gonic/gin"
)

func GetRoutes() {
	var port string = ":8000"

	var route = utils_functions.GetRouteUrls()

	var router *gin.Engine = utils_functions.GetGinRouter()

	routes_mails.MailRouteBasic(router, route.Mail.Basic)

	router.Run(port)
}
