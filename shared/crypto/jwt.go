package crypto

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const issuer = "stein.systems"

type CreateJWT func(subject string) (string, error)
type VerifyJWT func(token string) (jwt.MapClaims, error)

func NewCreateJWT(secret string) CreateJWT {
	return func(subject string) (string, error) {
		now := time.Now().UTC()
		claims := &jwt.StandardClaims{
			Issuer:    issuer,
			Audience:  issuer,
			Subject:   subject,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(24 * 30 * time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(secret))
	}
}

func NewVerifyJWT(secret string) VerifyJWT {
	return func(token string) (jwt.MapClaims, error) {
		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil {
			return map[string]interface{}{}, err
		}

		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			return claims, nil
		}
		return map[string]interface{}{}, errors.New("failed to validate jwt")
	}
}
