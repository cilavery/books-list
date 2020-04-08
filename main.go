package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `jsong:author`
	Year string `json:year`
}

var books []Book
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
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	params := mux.Vars(r)

	err := db.QueryRow("select * from books where id=$1", params["id"]).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	log.Fatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 RETURNING id;", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {

}