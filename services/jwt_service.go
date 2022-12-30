package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CustomClaims 自定义签名结构
type CustomClaims struct {
	jwt.RegisteredClaims
	UserName string `json:"user_name"`
	ID       int64  `json:"id"`
}

// CustomSecret 用于签名的盐值
var CustomSecret = []byte("ZPhwPk3vwzQ3Lc$")
var ExpireAt = time.Now().Add(time.Hour * 24 * 7)

// GetToken 生成Token
func GetToken(userName string, ID int64) (string, error) {
	// 创建 Claims
	claims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpireAt), // 过期时间
			Issuer:    "sqy",                        // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userName,
		ID,
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("token无效")
}
