package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sean-tech/gokit/foundation"
	"strconv"
	"sync"
	"time"
)

const (
	// jwt token
	STATUS_CODE_AUTH_CHECK_TOKEN_FAILED    = 801
	STATUS_CODE_AUTH_CHECK_TOKEN_TIMEOUT   = 802
	STATUS_CODE_AUTH_TOKEN_GENERATE_FAILED = 803
	STATUS_CODE_AUTH_TYPE_ERROR                = 804
	// jwt token
	STATUS_MSG_AUTH_CHECK_TOKEN_FAILED    = "Token鉴权失败"
	STATUS_MSG_AUTH_CHECK_TOKEN_TIMEOUT   = "Token已过期"
	STATUS_MSG_AUTH_TOKEN_GENERATE_FAILED = "Token生成失败"
	STATUS_MSG_AUTH_TYPE_ERROR            = "Token校验类型错误"
)

type TokenInfo struct {
	UserId uint64 			`json:"userId"`
	UserName string 		`json:"userName"`
	IsAdministrotor bool 	`json:"isAdministrotor"`
	jwt.StandardClaims
}

type ITokenManager interface {
	GenerateToken(userId uint64, userName string, isAdministrotor bool, JwtSecret string, JwtIssuer string, JwtExpiresTime time.Duration) (string, error)
	ParseToken(token string, JwtSecret string, JwtIssuer string) (*TokenInfo, error)
	CheckToken(token string, JwtSecret string, JwtIssuer string) error
}

var (
	_tokenManagerOnce sync.Once
	_tokenManager     ITokenManager
)
/**
 * 获取jwt管理实例
 */
func GetTokenManager() ITokenManager {
	_tokenManagerOnce.Do(func() {
		_tokenManager = &tokenManagerImpl{}
	})
	return _tokenManager
}

/**
 * jwt实现
 */
type tokenManagerImpl struct{}

/**
 * 生成token
 */
func (this *tokenManagerImpl) GenerateToken(userId uint64, userName string, isAdministrotor bool, JwtSecret string, JwtIssuer string, JwtExpiresTime time.Duration) (string, error) {
	expireTime := time.Now().Add(JwtExpiresTime)
	iat := time.Now().Unix()
	jti, _ := foundation.GenerateId(1)
	c := TokenInfo{
		UserId:			userId,
		UserName:       userName,
		IsAdministrotor:isAdministrotor,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    JwtIssuer,
			Id:strconv.FormatInt(jti, 10),
			IssuedAt:iat,
			NotBefore: iat,
			Subject:"client",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenClaims.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", foundation.NewError(STATUS_CODE_AUTH_TOKEN_GENERATE_FAILED, STATUS_MSG_AUTH_TOKEN_GENERATE_FAILED)
	}
	return token, nil
}

/**
 * 解析token
 */
func (this *tokenManagerImpl) ParseToken(token string, JwtSecret string, JwtIssuer string) (*TokenInfo, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, foundation.NewError(STATUS_CODE_AUTH_CHECK_TOKEN_FAILED, STATUS_MSG_AUTH_CHECK_TOKEN_FAILED)
	}
	if !tokenClaims.Valid {
		return nil, foundation.NewError(STATUS_CODE_AUTH_CHECK_TOKEN_FAILED, STATUS_MSG_AUTH_CHECK_TOKEN_FAILED)
	}
	claims, ok := tokenClaims.Claims.(*TokenInfo)
	if !ok {
		return nil, foundation.NewError(STATUS_CODE_AUTH_TYPE_ERROR, STATUS_MSG_AUTH_TYPE_ERROR)
	}
	if claims.Issuer != JwtIssuer {
		return nil, foundation.NewError(STATUS_CODE_AUTH_CHECK_TOKEN_FAILED, STATUS_MSG_AUTH_CHECK_TOKEN_FAILED)
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, foundation.NewError(STATUS_CODE_AUTH_CHECK_TOKEN_TIMEOUT, STATUS_MSG_AUTH_CHECK_TOKEN_TIMEOUT)
	}
	return claims, nil
}

func (this *tokenManagerImpl) CheckToken(token string, JwtSecret string, JwtIssuer string) error {
	if _, err := this.ParseToken(token, JwtSecret, JwtIssuer); err != nil {
		return err
	}
	return nil
}

