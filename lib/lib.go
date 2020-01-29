package lib

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTKey defines the token key
var JWTKey = "chrisju"

// GenerateJWT generate token for user
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sub": username,
	})

	return token.SignedString([]byte(JWTKey))
}

// CheckJWT check whether the jwt is valid and if it is in the invalid database
func CheckJWT(jwtString string) string {

	return validateToken(jwtString)
}

// validateToken check the format of token
func validateToken(jwtString string) string {
	// validate jwt
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expired := claims["exp"]
		if expired == nil {
			return ""
		}
		exp := int64(expired.(float64))
		if time.Now().Unix() > exp {
			return ""
		}
		return claims["sub"].(string)
	}
	return ""

}

// StringToMd5 Transfer string to Md5
func StringToMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// VerifyUsernameFormat method
func VerifyUsernameFormat(username string) bool {
	pattern := `^[a-z0-9_-]{6,20}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}
