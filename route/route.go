package route

import "github.com/gin-gonic/gin"

func BindRoute(server *gin.Engine) {
	server.POST("/send", sendEmailHandler())
}
