package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var Secret = initSecret()

func initSecret() []byte {
	return []byte(os.Getenv("SECRET"))
}

// GetParam ce face
func GetParam(asd map[string][]string, element string) (string, error) {
	param := asd[element]
	if len(param) > 0 {
		return param[0], nil
	}
	return "", &ErrorString{Str: "Nu ai parametru"}
}

//ParseJWT stfu
func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
}

//GetClaims stfu
func GetClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := ParseJWT(tokenString)
	claims, _ := token.Claims.(jwt.MapClaims)
	if err != nil {
		return nil, &ErrorString{Str: "Error parsing token."}
	}
	return claims, nil
}

//ErrorJSONstring encodes an error in JSON
func ErrorJSONstring(err *ErrorString) ([]byte, error) {
	final := &ErrorJSON{Message: err.Str, Success: false}
	return json.Marshal(final)
}
