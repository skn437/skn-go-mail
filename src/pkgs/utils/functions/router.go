package utils_functions

import "github.com/gin-gonic/gin"

var router *gin.Engine = gin.Default()

func GetGinRouter() *gin.Engine {
	return router
}
