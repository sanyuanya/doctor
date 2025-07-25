package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sanyuanya/doctor/config"
)

// JWTClaims JWT 载荷结构
type JWTClaims struct {
	UserID      uint   `json:"user_id"`
	OpenID      string `json:"open_id"`
	PhoneNumber string `json:"phone_number"`
	DefaultRole string `json:"default_role"`
	ActiveRole  string `json:"active_role"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成 JWT token
func GenerateJWT(userID uint, openID, phoneNumber, defaultRole, activeRole string) (string, error) {
	// 设置过期时间（7天）
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 创建 claims
	claims := &JWTClaims{
		UserID:      userID,
		OpenID:      openID,
		PhoneNumber: phoneNumber,
		DefaultRole: defaultRole,
		ActiveRole:  activeRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "doctor-app",
			Subject:   openID,
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析 JWT token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token 并提取 claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
