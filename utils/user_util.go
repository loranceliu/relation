package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gin/entity/vo"
	"github.com/dgrijalva/jwt-go"
	"io"
	"math/rand"
	"time"
)

const Secret = "aaabbbccc"

func GenerateToken(vo *vo.UserInfoVO) (string, int, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	exp := 720

	expirationTime := time.Now().Add(time.Duration(exp) * time.Minute)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = vo.UserId
	claims["name"] = vo.Name
	claims["email"] = vo.Email
	claims["username"] = vo.Username
	claims["exp"] = expirationTime.Unix() // 设置过期时间

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

func GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	saltLength := 10

	saltBytes := make([]byte, saltLength)
	for i := range saltBytes {
		saltBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(saltBytes)
}

func HashMD5(input string) string {
	hash := md5.New()
	_, err := io.WriteString(hash, input)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
