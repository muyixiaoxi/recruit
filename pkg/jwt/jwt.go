package jwt

import (
	"errors"
	"strconv"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var mySecret = []byte("夏天夏天悄悄过去")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

const TokenExpireDuration = time.Hour * 24 * 365

// GenToken 生成access token 和 refresh token
func GenToken(userID uint) (token string, err error) {
	// 创建一个我们自己的声明
	c := jwt.StandardClaims{
		Id:        strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
		Issuer:    "mosi",
	}

	// 使用指定的签名方法创建签名对象
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}
