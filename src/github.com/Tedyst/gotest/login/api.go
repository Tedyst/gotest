package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tedyst/gotest/util"
)

//Handler stfu
func Handler(w http.ResponseWriter, r *http.Request) {
	jwt, err := Authenticate(r)
	if jwt != "" {
		fmt.Fprintf(w, jwt)
		return
	}
	if err != nil {
		asd := err.(*util.ErrorString)
		str, _ := util.ErrorJSONstring(asd)
		fmt.Fprintf(w, string(str))
		return
	}
	asd := &util.Response{Message: "Logged in", Success: true}
	str, _ := json.Marshal(asd)
	fmt.Fprintf(w, string(str))
}
