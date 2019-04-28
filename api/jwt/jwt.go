package jwt

import (
	"github.com/apex/log"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var signingKey = os.Getenv("SIGNING_KEY")

type Helper struct {
	SigningKey string
}

func NewHelper() *Helper {
	return &Helper{
		SigningKey: signingKey,
	}
}

func (jwtHelper *Helper) Encode(claim jwt.Claims) string {
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claim).
		SignedString(signingKey)
	if err != nil {
		log.WithError(err).Fatalf("Unable to encode JWT %+v", claim)
	}
	return signedString
}

func (jwtHelper *Helper) Decode(token string) jwt.Claims {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return signingKey, nil
	})
	if err != nil {
		log.WithError(err).Fatalf("Unable to decode JWT %+v", token)
	}
	return claims
}
