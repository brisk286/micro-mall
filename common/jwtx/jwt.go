package jwtx

import "github.com/golang-jwt/jwt"

// GetToken 生成token
// 密钥 签发时间 过期时间  responseId
func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	// 过期时间 = 签发时间 + 期限
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	// 使用HS256算法加密
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
