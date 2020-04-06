package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `jsong:author`
	Year string `json:year`
}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{1, "Book Title One", "Author One", "1999"},
		Book{2, "Book Title Two", "Author Two", "2000"},
		Book{3, "Book Title Three", "Author Three", "2001"},
		Book{4, "Book Title Four", "Author Four", "2002"},
		Book{5, "Book Title Five", "Author Five", "2003"},
		Book{6, "Book Title Six", "Author Six", "2004"},
	)

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all book is called")
	params := mux.Vars(r)

	//get type of
	//log.Println(reflect.TypeOf(params["id"]))

	int, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == int {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
	var nb Book
	json.NewDecoder(r.Body).Decode(&nb)
	books = append(books, nb)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book is called")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book is called")
	params := mux.Vars(r)
	int, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == int {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}