package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	r := mux.NewRouter()

	r.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
	fmt.Println("Server is listening on PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}