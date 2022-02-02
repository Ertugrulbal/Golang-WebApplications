package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/set", set)
	// http.HandleFunc("/get", get)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		sessionid := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    sessionid.String(),
			HttpOnly: true,
			// Secure: true,
			Path: "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

// func set(w http.ResponseWriter, r *http.Request) {

// }

// func get(w http.ResponseWriter, r *http.Request) {

// }
