package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateECPrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func Sign(userId uint, privateKey *ecdsa.PrivateKey) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": userId,
		"exp":    expirationTime.Unix(),
	})
	return token.SignedString(privateKey)
}

func Validate(s string, publicKey *ecdsa.PublicKey) (uint, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["userId"].(float64)
		if ok {
			return uint(userId), nil
		}
		return 0, errors.New("failed type conversion")
	} else {
		return 0, errors.New("failed token validation")
	}
}
