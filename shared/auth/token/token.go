package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (any, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not StandardClaims")
	}

	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid: %v", err)
	}

	return clm.Subject, nil
}
