package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

const lockStream = "st=lDEFABCkVWXYZabc89LMmGH012345uvdefIJK6NOPyzghijQRSTUwx7nopqr"
const keyPassword = "szmgszmgszmgszmg"

// SimpleEncrypt /**
func SimpleEncrypt(textSteam string) (encodeStr string) {
	var k int
	var j int
	stream := []byte(lockStream)

	lockLen := len(stream)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	lockCount := ran.Intn(lockLen)

	randomLock := string(stream[lockCount])

	password := func(str string) string {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}(keyPassword + randomLock)

	textSteam = base64.StdEncoding.EncodeToString([]byte(textSteam))

	for _, ts := range textSteam {
		if k == len(password) {
			k = 0
		}
		j = bytes.IndexFunc(stream, func(r rune) bool {
			return r == ts
		})
		j = (j + lockCount + int(password[k])) % (lockLen)
		encodeStr = encodeStr + string(stream[j])
		k++
	}
	return encodeStr + randomLock
}

// SimpleDecrypt /**
func SimpleDecrypt(textSteam string) (decodeStr string) {
	var k int
	var j int
	stream := []byte(lockStream)

	lockLen := len(stream)

	// 截取随机密锁值
	textLen := len(textSteam)
	randomLock := textSteam[textLen-1]
	// 获得随机密码值的位置
	lockCount := bytes.IndexFunc(stream, func(r rune) bool {
		return r == rune(randomLock)
	})
	password := func(str string) string {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}(keyPassword + string(randomLock))

	// 开始对字符串解密
	textSteam = textSteam[:textLen-1]
	for _, ts := range textSteam {
		if k == len(password) {
			k = 0
		}
		j = bytes.IndexFunc(stream, func(r rune) bool {
			return r == ts
		})

		j = j - lockCount - int(password[k])

		for {
			if j >= 0 {
				break
			}
			j = j + lockLen
		}
		decodeStr += string(lockStream[j])
		k++
	}
	decodeByte, _ := base64.StdEncoding.DecodeString(decodeStr)
	decodeStr = string(decodeByte)
	return
}
