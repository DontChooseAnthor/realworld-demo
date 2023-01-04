package middleware

import (
	"fmt"
	"testing"
)

var Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjQ2NDYzNzI0M0BxcS5jb20iLCJwYXNzd29yZCI6IndzYXdzZHdzIiwiZXhwIjoxNjYyMzAzMjc2fQ.I9tQ_bmxCDm4kayIdnROeptSMCbSEvR4eKqtgfYlceE"

func TestJWTGenerate(t *testing.T) {
	token, err := JWTGenerate("464637243@qq.com", "wsawsdws")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}

func TestJWTAuth(t *testing.T) {
	claims, err := JWTAuth(Token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims.Email, claims.Password)
}
