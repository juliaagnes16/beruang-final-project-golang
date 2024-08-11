package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		err = jwt.ErrSignatureInvalid
		return
	}
	return
}