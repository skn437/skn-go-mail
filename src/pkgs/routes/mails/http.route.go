package mails

import "github.com/gin-gonic/gin"

func HttpMailRoute(router *gin.Engine, url string) {
	router.POST(url, func(ctx *gin.Context) {
		var data *HttpMailDataType = &HttpMailDataType{}

		var err error = ctx.BindJSON(&data)

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, gin.H{
			"post-title":       data.Title,
			"post-description": data.Description,
		})
	})

	//* Query String Implementation
	router.GET(url, func(ctx *gin.Context) {
		var name string = ctx.Query("name")
		var id string = ctx.Query("id")

		ctx.JSON(200, gin.H{
			"name": name,
			"id":   id,
		})
	})

	//* Param Implementation
	router.GET(url+"/:id", func(ctx *gin.Context) {
		var id string = ctx.Param("id")

		ctx.JSON(200, gin.H{
			"id":    id,
			"title": "URL Param",
		})
	})
}
