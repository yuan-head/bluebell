package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2 // Token有效期，单位为秒

var mySecret = []byte("my_secret_key") // JWT密钥

type MyClaims struct {
	UserID             int64  `json:"user_id"`  // 用户ID
	Username           string `json:"username"` // 用户名
	jwt.StandardClaims        // 内置的标准声明
}

// GenToken 生成JWT Token
func GenToken(userID int64, username string) (string, error) {
	// 创建一个新的声明的数据
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 设置Token的过期时间
			Issuer:    "wainlin-bluebell",                         // 设置Token的签发时间
		},
	}
	// 使用HS256算法创建一个新的Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用密钥签名Token并返回字符串形式的Token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析Token字符串，返回一个Token对象
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
