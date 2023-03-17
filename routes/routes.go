package routes

import (
	"net/http"

	"server/controllers"
	"server/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "route success")
	})
	r.POST("/signUp", controllers.SignUpHandler)
	r.GET("/fuck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "fuck",
			"code": 200,
		})
	})
	return r
}
