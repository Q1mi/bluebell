package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

var mySecret = []byte("夏天夏天悄悄过去")

func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

const TokenExpireDuration = time.Second

// GenToken 生成access token 和 refresh token
func GenToken(userID uint64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发人
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newToken string, err error) {
	// refresh token无效直接返回
	_, err = jwt.Parse(rToken, keyFunc)
	if err != nil {
		return "", err
	}

	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		fmt.Println(1, claims)

		return GenToken(claims.UserID)
	}
	return "", err
}
