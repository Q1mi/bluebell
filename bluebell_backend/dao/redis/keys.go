package redis

/*
	Redis Key
*/

const (
	KeyArticleInfoHashFormat = "bluebell:article:%d"
	KeyArticleTimeZSet       = "bluebell:article:time"
	KeyArticleScoreZSet      = "bluebell:article:score"
	KeyArticleVotedSetFormat = "bluebell:article:voted:%d"

	KeyCommunitySetFormat = "bluebell:community:%s"
)
