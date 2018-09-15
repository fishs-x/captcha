package router

import (
	"insur-box/src/controller"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// 使用middleware
	//authorized := router.Group("/api/v1",middleware.AuthRequired)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/get/captcha", controller.GetCaptcha)
		v1.POST("/verify/captcha", controller.VerifyCaptcha)
	}

}
