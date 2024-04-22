package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"bms-server/global"
	"bms-server/model/members/request"
)

type JWTMember struct {
	SigningKey []byte
}

var (
	MemberTokenExpired     = errors.New("token is expired")
	MemberTokenNotValidYet = errors.New("token not active yet")
	MemberTokenMalformed   = errors.New("that's not even a token")
	MemberTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWTMember() *JWTMember {
	return &JWTMember{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWTMember) CreateClaims(baseClaims request.BaseMemberClaims) request.CustomMemberClaims {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := request.CustomMemberClaims{
		BaseMemberClaims: baseClaims,
		BufferTime:       int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"BMS"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    "bmsCool",                                 // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWTMember) CreateToken(claims request.CustomMemberClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWTMember) CreateTokenByOldToken(oldToken string, claims request.CustomMemberClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("BMS_JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWTMember) ParseToken(tokenString string) (*request.CustomMemberClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomMemberClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, MemberTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, MemberTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, MemberTokenNotValidYet
			} else {
				return nil, MemberTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomMemberClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, MemberTokenInvalid

	} else {
		return nil, MemberTokenInvalid
	}
}
