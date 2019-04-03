package api

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = initSecret()

//GetJWTKey stfu
func GetJWTKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}

//ParseJWT stfu
func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}

func initSecret() []byte {
	return []byte(os.Getenv("SECRET"))
}

//GetClaims stfu
func GetClaims(tokenString string) (jwt.MapClaims, error) {
	tokenA, err := ParseJWT(tokenString)
	if err != nil {
		return nil, &ErrorString{Str: "Error validating token."}
	}
	claims, ok := tokenA.Claims.(jwt.MapClaims)
	if claims["date"] != nil {
		date := time.Now().Unix() - int64(claims["date"].(float64))
		// One day
		if date > 86400 {
			return nil, &ErrorString{Str: "Token expired."}
		}
	}
	if ok && tokenA.Valid {
		return claims, nil
	}
	return nil, &ErrorString{Str: "Invalid token."}
}

//CreateJWT creeaza un JWT token
func CreateJWT(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user,
		"date": int(time.Now().Unix()),
	})
	return token.SignedString(secret)
}

func Authenticate(user string, password string) bool {
	if user == "Tedyst" && password == "Tedyst" {
		return true
	}
	return false
}
