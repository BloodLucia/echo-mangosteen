package jwt

import (
	"echo-mangosteen/pkg/config"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

type MyCustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func New(conf *config.Config) *JWT {
	return &JWT{key: []byte(conf.JWT.Key)}
}

func (j *JWT) BuildToken(userId string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        userId,
			Audience:  []string{},
		},
	})
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
