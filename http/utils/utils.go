package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtsecret = []byte(GetSecret())


type Claims struct {
	Name string `json:"username"`
	Email string `json:"email"`
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GetSecret() string{
	secret := os.Getenv("JWT_SECRET")

	if secret =="" {
		return "hahanosecret"
	}
	return secret
}

func CreateToken(userId, name, email, password string) (string, error){

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Name : name,
		Email : email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject: userId,
		},
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtsecret)

	if err!=nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error){
		return jwtsecret, nil
	})

	if err!=nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}