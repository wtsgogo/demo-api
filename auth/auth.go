package auth

import (
	"demo-api/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Password string
}

var jwtSecret = []byte(config.AppJwtSecret)

func GenerateToken(password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(2 * time.Hour)
	claims := Claims{
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
