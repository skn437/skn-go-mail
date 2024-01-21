package main

import (
	"skn-go-mail/src/pkgs/middlewares"

	"github.com/gin-gonic/gin"
)

//"log"
//"skn-go-mail/src/pkgs/routes"
//"skn-go-mail/src/pkgs/utils"

func main() {
	//routes.GetRoutes()

	//var base string = utils.ViperEnvVariable("ROUTES.MAIL.BASIC")

	//log.Fatalf("Base: %s", base)
	var router *gin.Engine = gin.New()

	//* Middlewares run sequentially i.e. the first one will run first, second one will run second and so on
	router.Use(gin.Logger(), middlewares.CommonMiddleware) //* This ensures middleware for all routes

	var root *gin.RouterGroup = router.Group("/api", middlewares.CsrfMiddleware("csrf"), middlewares.AddHeader("Wang So!"))
	{
		root.GET("/wang", func(ctx *gin.Context) {
			ctx.String(200, "Hello SKN!")
		})
	}

	var mail *gin.RouterGroup = router.Group("/mail")
	{
		mail.GET("/wang", func(ctx *gin.Context) {
			ctx.String(200, "You are good to go!")
		})
	}

	var err error = router.Run(":8000")

	if err != nil {
		panic(err)
	}
}
