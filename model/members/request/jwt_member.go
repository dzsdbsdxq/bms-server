package request

import "github.com/golang-jwt/jwt/v4"

// Custom claims structure
type CustomMemberClaims struct {
	BaseMemberClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseMemberClaims struct {
	UUID        string
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
