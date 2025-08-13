package utils

import (
	"time"

	"github.com/DiansSopandi/goride_be/pkg"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// jwtSecret := pkg.Cfg.Application.JwtSecret
	// SigningKey: jwtware.SigningKey{Key: []byte(pkg.Cfg.Application.SsoJwtSecret)},

	jwtSecret := pkg.Cfg.Application.JwtSecretKey
	signedToken, signErr := token.SignedString([]byte(jwtSecret))

	if signErr != nil {
		return "", signErr
	}

	return signedToken, nil
}
