package mails

import (
	"github.com/gin-gonic/gin"
)

type DataType struct {
	Name  string //* This fields should be public but in the 'post' request, request body can have lower case names
	Title string
}

func MailRouteBasic(router *gin.Engine, group string, url string) {
	var mail *gin.RouterGroup = router.Group(group)

	{
		mail.POST(url, func(ctx *gin.Context) {
			//var body io.ReadCloser = ctx.Request.Body
			//value, err := io.ReadAll(body)

			var data *DataType
			var err error = ctx.BindJSON(&data)

			if err != nil {
				panic(err)
			}

			ctx.JSON(200, gin.H{
				"message-name":  data.Name,
				"message-title": data.Title,
			})
		})

		mail.GET(url, func(ctx *gin.Context) {
			ctx.String(200, "Email Sent Successfully!")
		})
	}
}
