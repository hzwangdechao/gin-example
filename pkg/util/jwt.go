package util

import (
	"gin-example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Clamis struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 根据用户名和密码生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 3)
	clamis := Clamis{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-example",
		},
	}
	tokenClamis := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	token, err := tokenClamis.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Clamis, error) {
	tokenClamis, err := jwt.ParseWithClaims(token, &Clamis{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, err

	})

	if tokenClamis != nil {
		if clamis, ok := tokenClamis.Claims.(*Clamis); ok && tokenClamis.Valid {
			return clamis, nil
		}
	}
	return nil, err

}
