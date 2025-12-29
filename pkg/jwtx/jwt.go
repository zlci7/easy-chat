package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
)

// GetToken 生成 JWT Token
// secretKey: 密钥 (来自配置文件)
// iat: 当前时间戳 (Seconds)
// seconds: 过期时间 (Seconds)
// uid: 用户ID
func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	// 这个 key "uid" 非常重要，后续在 API 网关解析时会用到
	claims["uid"] = uid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
