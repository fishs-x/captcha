package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context)  {
	c.JSON(200, gin.H{"OK": "No"})
	// 终止
	c.Abort()
}
