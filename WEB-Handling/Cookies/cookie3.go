package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Mini Cookie App (Counter)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("countercookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "countercookie",
			Value: "0",
			Path:  "/",
		}
	}

	counter, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	counter++
	cookie.Value = strconv.Itoa(counter)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)

	fmt.Fprintln(w, "Cookieleri kontrol et!")
}
