package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	email := "test@example.com"
	tokenString, err := GenerateJWT(email)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestValidateJWTToken(t *testing.T) {
	email := "test@example.com"
	tokenString, err := GenerateJWT(email)
	assert.NoError(t, err)

	validEmail, err := ValidateJWTToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, email, validEmail)
}

func TestValidateExpiredJWTToken(t *testing.T) {
	// Create an expired token
	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &JWTClaim{
		Email: "test@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	assert.NoError(t, err)

	_, err = ValidateJWTToken(tokenString)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired")
}

func TestGetJWTKey(t *testing.T) {
	key := GetJWTKey()
	assert.NotEmpty(t, key)
}
