package util

import (
	"ginDemo/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
*

总的来说 jwt的任务就是，生成token和解析token
*/
var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 这种写法，username是否意味着同样是string的写法
func GenerateToken(username, password string) (string, error) {
	// 生成token 的关键点
	// 加密内容，明文内容，加密的算法
	now := time.Now()
	// 过期时间
	expireTime := now.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "LuckKun-Gin-Blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 利用secret再进行一次加密
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 这里不需要加密算法，是不是因为claims里面已经有了
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
