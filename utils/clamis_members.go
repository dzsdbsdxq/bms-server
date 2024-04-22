package utils

import (
	"bms-server/global"
	membersReq "bms-server/model/members/request"
	"github.com/gin-gonic/gin"
	"net"
)

func ClearMemberToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func SetMemberToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetMemberToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

func GetMemberClaims(c *gin.Context) (*membersReq.CustomMemberClaims, error) {
	token := GetToken(c)
	j := NewJWTMember()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetMemberID 从Gin的Context中获取从jwt解析出来的用户ID
func GetMemberID(c *gin.Context) uint {
	if claims, exists := c.Get("memberClaims"); !exists {
		if cl, err := GetMemberClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseMemberClaims.ID
		}
	} else {
		waitUse := claims.(*membersReq.CustomMemberClaims)
		return waitUse.BaseMemberClaims.ID
	}
}

// GetMemberUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetMemberUuid(c *gin.Context) string {
	if claims, exists := c.Get("memberClaims"); !exists {
		if cl, err := GetMemberClaims(c); err != nil {
			return ""
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*membersReq.CustomMemberClaims)
		return waitUse.UUID
	}
}

// GetMemberInfo 从Gin的Context中获取从jwt解析出来的用户信息
func GetMemberInfo(c *gin.Context) *membersReq.CustomMemberClaims {
	if claims, exists := c.Get("memberClaims"); !exists {
		if cl, err := GetMemberClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*membersReq.CustomMemberClaims)
		return waitUse
	}
}

// GetMemberUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetMemberUserName(c *gin.Context) string {
	if claims, exists := c.Get("memberClaims"); !exists {
		if cl, err := GetMemberClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*membersReq.CustomMemberClaims)
		return waitUse.Username
	}
}
