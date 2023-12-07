package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/sirupsen/logrus"
)

var jwtKey []byte

func init() {

	// Load configuration during package initialization
	config, err := initializers.LoadConfig(".")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load configuration")
	}

	jwtKey = []byte(config.JWTTokenSecret)
}

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateJWTToken(signedToken string) (res string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token expired")
		return "", err
	}
	return claims.Email, nil
}

// GetJWTKey returns the JWT key
func GetJWTKey() []byte {
	return jwtKey
}
