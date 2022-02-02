package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cihanozhan/controllers"
	"github.com/cihanozhan/driver"
	"github.com/cihanozhan/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB
var books []models.Book

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	// Router ya da Endpoint
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	fmt.Println("Server has started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
