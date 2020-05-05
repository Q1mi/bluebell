package redis

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds = 7 * 24 * 3600
	VoteScore        = 432
	PostPerAge       = 20
)

// PostVote 为帖子投票
func PostVoteUp(postID, userID string, isDown bool) (err error) {
	// 1. 取帖子发布时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		// 不允许投票了
		return
	}
	upKey := KeyPostVotedUpSetPrefix + postID
	downKey := KeyPostVotedDownSetPrefix + postID
	key := upKey
	reverse := float64(1)
	if isDown {
		key = downKey
		reverse = -1
	}
	// 记录投票相关数据
	if !client.SIsMember(upKey, userID).Val() && !client.SIsMember(downKey, userID).Val() {
		client.SAdd(key, userID)
		client.ZIncrBy(KeyPostScoreZSet, VoteScore*reverse, postID)
		client.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	}
	return
}

func PostVoteReverse(postID, userID string, isDown bool) (err error) {
	// 1. 取帖子发布时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		// 不允许投票了
		return
	}
	sKey := KeyPostVotedUpSetPrefix + postID
	dKey := KeyPostVotedDownSetPrefix + postID
	if isDown {
		sKey, dKey = dKey, sKey
	}
	client.SMove(sKey, dKey, userID)
	// 记录投票相关数据

	if !client.SIsMember(sKey, userID).Val() {
		client.ZIncrBy(KeyPostScoreZSet, VoteScore, postID)
		client.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	}
	return
}

// CreatePost 使用hash存储帖子信息
func CreatePost(postID, userID, title, summary string) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedSetPrefix + postID
	pipeline := client.Pipeline()
	pipeline.SAdd(votedKey, userID)
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds)
	fields := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}
	pipeline.HMSet(KeyPostInfoHashPrefix+postID, fields)
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})
	_, err = pipeline.Exec()
	return
}

// GetPost 从key中分页取出帖子
func GetPost(key string, page int64) []map[string]string {
	start := (page - 1) * PostPerAge
	end := start + PostPerAge - 1
	ids := client.ZRevRange(key, start, end).Val()
	postList := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		postData := client.HGetAll(KeyPostInfoHashPrefix + id).Val()
		postData["id"] = id
		postList = append(postList, postData)
	}
	return postList
}

// GetCommunityPost 分社区根据发帖时间或者分数取出分页的帖子
func GetCommunityPost(communityName, orderKey string, page int64) []map[string]string {
	key := orderKey + communityName // 创建缓存键

	if client.Exists(key).Val() < 1 {
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, KeyCommunityPostSetPrefix+communityName, orderKey)
		client.Expire(key, 60*time.Second)
	}
	return GetPost(key, page)
}
