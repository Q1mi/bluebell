package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/utils"
	"errors"
	"fmt"
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

func RegisterHandler(c *gin.Context) {
	// 1.获取请求参数
	var fo models.RegisterForm
	if err := c.ShouldBind(&fo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeSuccess,
			"msg":  GetMsg(CodeSuccess),
		})
		return
	}
	// 2.校验数据有效性
	if ok, errMsg := fo.Validate(); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeInvalidParams,
			"msg":  errMsg,
		})
		return
	}
	// 3.注册用户
	err := mysql.Register(&models.User{
		UserName: fo.UserName,
		Password: fo.Password,
	})
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	if ok, errMsg := u.LoginValidate(); !ok {
		fmt.Println(errMsg)
		ResponseErrorWithMsg(c, CodeInvalidParams, errMsg)
		return
	}
	if err := mysql.Login(&u); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, nil)
}

func Send(value interface{}) {
	ch <- value
}
