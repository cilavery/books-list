package controllers

import (
	"books-list/models"
	"books-list/repository/book"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])
		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		//add error handler for missing book data author, title, year here
		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		// add error handler for server error here
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		// add error handler if see if book id exists?
		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All Fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		// add error handling for server error
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsUpdated == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		bookRepo := bookRepository.BookRepository{}
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		rd, err := bookRepo.RemoveBook(db, id)

		// add error handling
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rd == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rd)

	}
}