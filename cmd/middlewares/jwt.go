package middlewares

import (
	"errors"
	"time"

	"github.com/go-template-boilerplate/cmd/models"
	"github.com/go-template-boilerplate/generated"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *generated.User, env *models.EnvModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expired":  time.Now().Add(time.Hour).Unix(),
	})
	secret := []byte(env.JwtSecret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string, env *models.EnvModel) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(env.JwtSecret), nil
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

func GenerateRefreshToken(user *generated.User, env *models.EnvModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expired":  time.Now().Add(time.Hour * 7 * 24).Unix(),
	})
	secret := []byte(env.JwtSecret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GeneratedAccessAndRefreshTokens(user *generated.User, env *models.EnvModel) (string, string, error) {
	accessToken, err := GenerateToken(user, env)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenerateRefreshToken(user, env)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
