package controller

import (
	"bluebell_backend/models"
	"bluebell_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.Name == "q1mi" && user.Password == "q1mi123" {
		// 生成Token
		tokenString, _ := utils.GenToken(user.Name)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
