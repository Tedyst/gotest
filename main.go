package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Tedyst/gotest/login"
)

func main() {
	http.HandleFunc("/login", login.Handler)
	log.Println("Listening on port 3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Println("sal", err)
	}
}

func site(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("static", "index.html")
	print(lp)
	http.ServeFile(w, r, lp)
}
