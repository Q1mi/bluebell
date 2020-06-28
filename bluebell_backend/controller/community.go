package controller

import (
	"bluebell_backend/dao/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

// CommunityDetailHandler 社区详情
func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")
	communityList, err := mysql.GetCommunityByID(communityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
