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
	err := c.BindJSON(&user)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 校验用户名和密码是否正确
	if user.UserName == "q1mi" && user.Password == "q1mi123" {
		// 生成Token
		tokenString, _ := utils.GenToken(user.UserName)
		ResponseSuccess(c, tokenString)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeSuccess,
			"msg":  GetMsg(CodeSuccess),
		})
		return
	}
	// check username & password
	if u.PasswordValid() {

	}
}
