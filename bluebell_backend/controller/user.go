package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/jwt"
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	// 1.获取请求参数 2.校验数据有效性
	var fo models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
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
		zap.L().Error("mysql.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := mysql.Login(&u); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 生成Token
	tokenString, _ := jwt.GenToken(u.UserID)
	ResponseSuccess(c, gin.H{
		"token":    tokenString,
		"userID":   u.UserID,
		"userName": u.UserName,
	})
}
