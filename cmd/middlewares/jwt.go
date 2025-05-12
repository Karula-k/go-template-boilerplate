package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/go-template-boilerplate/generated"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *generated.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expired":  time.Now().Add(time.Hour).Unix(),
	})
	secret := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(float64); ok {
			return int64(id), nil
		}
		return 0, errors.New("id not found")
	}
	return 0, errors.New("invalid token")
}

func GenerateRefreshToken(user *generated.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expired":  time.Now().Add(time.Hour * 7 * 24).Unix(),
	})
	secret := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GeneratedAccessAndRefreshTokens(user *generated.User) (string, string, error) {
	accessToken, err := GenerateToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
