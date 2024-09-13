package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"onycom/models"
	"time"
)

//jwt 생성, 검증 함수 만들기

var jwtKey = []byte("onycom_jwt_key")

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
	claims := &Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "onycom",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := &Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "onycom",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (uint, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		// 유효기간이 지났는지 확인
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, ErrTokenExpired
		}
		return 0, ErrInvalidToken
	}

	// 토큰이 유효하지 않으면
	if !token.Valid {
		return 0, ErrInvalidToken
	}

	// 만료 시간이 현재 시간보다 이전인지 체크 (추가 검증)
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	return claims.UserId, nil

}

func CheckRefreshTokenInDB(userID uint) error {
	var token models.Token
	err := DB.Where("user_id = ?", userID).First(&token).Error
	if err != nil {
		return err
	}

	if time.Since(token.CreatedAt) > 2*time.Hour {
		return errors.New("refresh token expired")
	}

	return nil
}
