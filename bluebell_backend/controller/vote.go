package controller

import (
	"bluebell_backend/dao/redis"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"direction"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction float64 `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}

func VoteHandler(c *gin.Context) {
	// 给哪个文章投什么票
	var vote VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
