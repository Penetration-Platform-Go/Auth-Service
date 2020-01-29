package lib

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// JWTKey defines the token key
var JWTKey = "chrisju"

// HashPassword return hash of password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// GenerateJWT generate token for user
func GenerateJWT(username string, isValid bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sub": username,
		"iat": time.Now().Unix(),
		//"jti":   GenerateUUID(),
		"valid": isValid,
	})

	return token.SignedString([]byte(JWTKey))
}

// GenerateUUID generate uuid
// TODO: return error
// func GenerateUUID() string {
// 	id, err := uuid.NewV1()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return id.String()
// }

// CompasePassword compare raw password with hashed one
func CompasePassword(raw, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw)) == nil
}

// CheckJWT check whether the jwt is valid and if it is in the invalid database
func CheckJWT(jwtString string) (bool, error) {
	isValid, err := ValidateToken(jwtString)
	if err != nil {
		return false, err
	}
	if !isValid {
		return false, nil
	}
	return true, nil
}

// ValidateToken check the format of token
func ValidateToken(jwtString string) (bool, error) {
	// validate jwt
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// parse time from jwt
	var exp int64
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expired := claims["exp"]
		if expired == nil {
			return false, nil
		}
		exp = int64(expired.(float64))
		if time.Now().Unix() > exp {
			return false, nil
		}
	} else {
		return false, nil
	}
	return true, nil
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
