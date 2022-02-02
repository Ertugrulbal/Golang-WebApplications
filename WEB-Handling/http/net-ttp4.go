package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

type Book struct {
  ID     string json:id
  Title  string json:title
  Author string json:author
  Year   string json:year
}

var books []Book

func main() {

  router := mux.NewRouter()

  books = append(books,
    Book{ID: "1", Title: "C Programming Language", Author: "Dennis Ritchie", Year: "1978"},
    Book{ID: "2", Title: "The Go Programming Language", Author: "Brian Kernighan", Year: "2015"},
    Book{ID: "3", Title: "The Rust Programming Language", Author: "Carol Nichols", Year: "2018"},
    Book{ID: "4", Title: "The C# Programming Language", Author: "Anders Hejlsberg", Year: "2004"},
    Book{ID: "5", Title: "Programming Python", Author: "Mark Lutz", Year: "1996"},
  )

  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/books", addBook).Methods("POST")
  router.HandleFunc("/books", updateBook).Methods("PUT")
  router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r) // ?adada=blab&soyad=ozhan
  // fmt.Println(params)
  for _, book := range books {
    if book.ID == params["id"] {
      json.NewEncoder(w).Encode(&book)
    }
  }
}

func addBook(w http.ResponseWriter, r *http.Request) {
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  books = append(books, book)
  json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)

  for i, item := range books {
    if item.ID == book.ID {
      books[i] = book
    }
  }
  json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var book Book

  book.ID = params["id"]

  for i, item := range books {
    if item.ID == book.ID {
      books = append(books[:i], books[i+1:]...)
    }
  }

  json.NewEncoder(w).Encode(books)
}