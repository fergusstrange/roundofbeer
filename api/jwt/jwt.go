package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"os"
)

var signingKey = []byte(os.Getenv("SIGNING_KEY"))

type Helper struct {
	SigningKey []byte
}

type RoundToken struct {
	RoundUrl string `json:"round_url"`
	jwt.StandardClaims
}

func NewHelper() *Helper {
	return &Helper{
		SigningKey: signingKey,
	}
}

func (jwtHelper *Helper) Encode(token *RoundToken) string {
	signedString, err := jwt.
		NewWithClaims(jwt.SigningMethodHS512, token).
		SignedString(signingKey)
	errors.LogFatal(err)
	return signedString
}

func (jwtHelper *Helper) Decode(token string) (RoundToken, error) {
	var roundToken RoundToken
	_, err := jwt.ParseWithClaims(token, &roundToken, func(token *jwt.Token) (i interface{}, e error) {
		return signingKey, nil
	})
	return roundToken, err
}
