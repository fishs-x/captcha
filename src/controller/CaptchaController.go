package controller

import (
	"github.com/gin-gonic/gin"
	"insur-box/src/service"
	"net/http"
	"insur-box/src/utils"
	"insur-box/src/models"
)

/**
	获取图片验证码地址
 */

func GetCaptcha(c *gin.Context) {
	var style models.ImageStyle
	c.Bind(&style)
	if style.Width == 0 || style.Height == 0 {
		style.Width = 240
		style.Height = 80
	}
	url, verifyId := service.ImagesCode(style)
	c.JSON(http.StatusOK, utils.Response(0, "", map[string]string{"url": url, "verifyId": verifyId}))
}

/**
	验证输入验证码
 */
func VerifyCaptcha(c *gin.Context) {
	var verifyCode models.VerifyCode
	c.Bind(&verifyCode)
	isRight := service.VerifyImageCode(verifyCode)
	c.JSON(http.StatusOK, utils.Response(0, "", map[string]bool{"result": isRight}))
}
