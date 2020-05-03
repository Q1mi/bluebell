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

func RegisterHandler(c *gin.Context) {
	// 1.获取请求参数
	var fo models.RegisterForm
	if err := c.ShouldBind(&fo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeSuccess,
			"msg":  CodeSuccess.Msg(),
		})
		return
	}
	// 2.校验数据有效性
	if ok, errMsg := fo.Validate(); !ok {
		ResponseErrorWithMsg(c, CodeInvalidParams, errMsg)
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
		ResponseErrorWithMsg(c, CodeInvalidParams, errMsg)
		return
	}
	if err := mysql.Login(&u); err != nil {
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 生成Token
	tokenString, _ := utils.GenToken(u.UserID)
	ResponseSuccess(c, tokenString)
}
