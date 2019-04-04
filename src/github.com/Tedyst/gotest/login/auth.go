package login

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tedyst/gotest/util"

	"github.com/dgrijalva/jwt-go"
)

//GetJWTKey stfu
func GetJWTKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return util.Secret, nil
}

//ParseJWT stfu
func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return util.Secret, nil
	})
}

//ValidateJWT stfu
func ValidateJWT(tokenString string) error {
	tokenA, err := ParseJWT(tokenString)
	if err != nil {
		return &util.ErrorString{Str: "Error validating token."}
	}
	claims, ok := tokenA.Claims.(jwt.MapClaims)
	if claims["date"] != nil {
		date := time.Now().Unix() - int64(claims["date"].(float64))
		// One day
		if date > 86400 {
			return &util.ErrorString{Str: "Token expired."}
		}
	}
	if ok && tokenA.Valid {
		return nil
	}
	return &util.ErrorString{Str: "Invalid token."}
}

//CreateJWT creeaza un JWT token
func CreateJWT(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user,
		"date": int(time.Now().Unix()),
	})
	return token.SignedString(util.Secret)
}

func isValid(user string, password string) bool {
	if user == "Tedyst" && password == "Tedyst" {
		return true
	}
	return false
}

//Authenticate is authenticating a user
func Authenticate(r *http.Request) (string, error) {
	url := r.URL.Query()
	tokenString, _ := util.GetParam(url, "jwt")
	if len(tokenString) > 0 {
		err := ValidateJWT(tokenString)
		if err != nil {
			return "", err
		}
	} else {
		user, _ := util.GetParam(r.URL.Query(), "user")
		password, _ := util.GetParam(r.URL.Query(), "password")
		if !isValid(user, password) {
			return "", &util.ErrorString{Str: "Invalid creds."}
		}

		tokenString, err := CreateJWT(user)
		if err != nil {
			return "", &util.ErrorString{Str: "Error."}
		}
		//Hack ++
		return tokenString, nil
	}
	return "", nil
}
