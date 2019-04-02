package api

import (
	"fmt"
	"log"
	"net/http"
)

type paramError struct {
	prob string
}

func (e *paramError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

// GetParam ce face
func GetParam(asd map[string][]string, element string) (string, error) {
	param := asd[element]
	if len(param) > 0 {
		return param[0], nil
	}
	return "", &paramError{prob: "Nu ai parametru"}
}

// Handler Outputs "Test"
func Handler(w http.ResponseWriter, r *http.Request) {
	asd := r.URL.Query()
	param, err := GetParam(asd, "test")
	if err == nil {
		log.Println("test - " + param)
	} else {
		log.Println("test error - ", err)
	}
	fmt.Fprintf(w, "Test")
}
