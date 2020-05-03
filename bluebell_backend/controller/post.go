package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/logger"
	"bluebell_backend/models"
	"bluebell_backend/pkg/gen_id"
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// PostHandler 创建帖子
func PostHandler(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 生成帖子ID
	postID, err := gen_id.GetID()
	if err != nil {
		logger.Error("gen_id.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		logger.Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.PostID = postID
	post.AuthorId = userID

	// 创建帖子
	if err := mysql.CreatePost(&post); err != nil {
		logger.Error("mysql.CreatePost(&post) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子列表
func PostListHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)
}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	postID := c.Param("id")
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		logger.Error("mysql.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		logger.Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	post.AuthorName = user.UserName
	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		logger.Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	post.CommunityName = community.CommunityName
	ResponseSuccess(c, post)
}
