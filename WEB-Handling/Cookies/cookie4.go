package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/set", set)
  http.HandleFunc("/get", get)
  http.HandleFunc("/expire", expire)

  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, <h1><a href="/set">Set Cookie></a></h1>)
}

func set(w http.ResponseWriter, r *http.Request) {
  http.SetCookie(w, &http.Cookie{
    Name:  "cerez",
    Value: "Cerezin degeri",
    Path:  "/",
  })
  fmt.Fprintln(w, <h1><a href="/get">Get Cookie></a></h1>)
}

func get(w http.ResponseWriter, r *http.Request) {
  c, err := r.Cookie("cerez")
  if err != nil {
    http.Redirect(w, r, "/set", http.StatusSeeOther)
    return
  }
  fmt.Fprintf(w, <h1>Cookie : %v</h1><br><h2><a href="/expire">Expire</a></h2>, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
  c, err := r.Cookie("cerez")
  if err != nil {
    http.Redirect(w, r, "/set", http.StatusSeeOther)
    return
  }
  c.MaxAge = -1
  http.SetCookie(w, c)
  http.Redirect(w, r, "/", http.StatusSeeOther)
}