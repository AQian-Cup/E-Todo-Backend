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

func Sign(name string, privateKey *ecdsa.PrivateKey) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"name": name,
		"exp":  expirationTime.Unix(),
	})
	return token.SignedString(privateKey)
}

func Validate(s string, publicKey *ecdsa.PublicKey) (interface{}, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["name"], nil
	} else {
		return "", errors.New("failed token validation")
	}
}
