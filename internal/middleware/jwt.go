package middleware

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"realworld/internal/conf"
	"time"
)

var (
	pwdValidErr   = "jwt password valid error"
	emailValidErr = "jwt email valid error"
	jwtParseErr   = "jwt parse with claims error"
	jwtCreateErr  = "jwt create with claims error"
)

var signature = []byte(conf.Auth.GetAccessSecret(conf.Auth{
	AccessSecret: "",
	AccessExpire: nil,
}))

type MyCustomClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func JWTAuth(token string) (*MyCustomClaims, error) {
	if token == "" {
		log.Error(pwdValidErr)
		return nil, errors.New(pwdValidErr)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})
	if err != nil {
		log.Error(jwtParseErr)
		return nil, errors.New(jwtParseErr)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New(pwdValidErr)
}

func JWTGenerate(email, password string) (string, error) {
	if email == "" {
		log.Error(emailValidErr)
		return "", errors.New(emailValidErr)
	}

	if password == "" {
		log.Error(pwdValidErr)
		return "", errors.New(pwdValidErr)
	}

	claims := &MyCustomClaims{
		Email:    email,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signature)
	if err != nil {
		log.Error(jwtCreateErr)
		return "", errors.New(jwtCreateErr)
	}
	return ss, nil
}
