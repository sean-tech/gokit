package service

import (
	"fmt"
	"testing"
	"time"
)

type iTokenManager interface {
	GenerateToken(userId uint64, userName string, isAdministrotor bool, JwtSecret string, JwtIssuer string, JwtExpiresTime time.Duration) (string, error)
	CheckToken(token string, JwtSecret string, JwtIssuer string) error
}
func TestToken(t *testing.T) {
	var secret = "ahsjdadusba"
	var issuer = "sean.test"
	var tokenMgr iTokenManager = GetTokenManager()

	var expires = 30 * time.Second
	token, err := tokenMgr.GenerateToken(1230090123, "seantest1", false, secret, issuer, expires)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("token generate success!---" + token)

	if err := tokenMgr.CheckToken(token, secret, issuer); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("token check success!")
}