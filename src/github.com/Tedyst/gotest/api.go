package api

import (
	"fmt"
	"net/http"
)

// GetParam ce face
func GetParam(asd map[string][]string, element string) (string, error) {
	param := asd[element]
	if len(param) > 0 {
		return param[0], nil
	}
	return "", &ErrorString{Str: "Nu ai parametru"}
}

//Handler stfu
func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	tokenString, _ := GetParam(url, "jwt")
	if len(tokenString) > 0 {
		claims, err := GetClaims(tokenString)
		if err != nil {
			fmt.Fprintf(w, err.(*ErrorString).Str)
			println(err.(*ErrorString).Str)
			return
		}
		fmt.Fprintf(w, claims["name"].(string))
	} else {
		user, _ := GetParam(r.URL.Query(), "user")
		password, _ := GetParam(r.URL.Query(), "password")
		if !Authenticate(user, password) {
			fmt.Fprintf(w, "Invalid creds.")
			return
		}

		tokenString, err := CreateJWT("Tedyst")
		if err != nil {
			fmt.Fprintf(w, "Error.")
			return
		}
		fmt.Fprintf(w, tokenString)
		return
	}
	fmt.Fprintf(w, "\nLogged in boi.")
}
