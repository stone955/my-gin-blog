package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 根据用户名和密码生成 token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my-gin-blog",
		},
	}

	// 加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	token, err := tokenClaims.SignedString(setting.AppCfg.JwtSecretBytes)

	return token, err
}

// 解析 token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return setting.AppCfg.JwtSecretBytes, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
