package token

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		privateKey: privateKey,
		nowFunc:    time.Now,
	}
}

func (g *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := g.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: nowSec + int64(expire.Seconds()),
		IssuedAt:  nowSec,
		Issuer:    g.issuer,
		Subject:   accountID,
	})

	return tkn.SignedString(g.privateKey)
}
