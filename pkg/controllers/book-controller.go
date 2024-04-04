package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kiyuu10/go-bookstore/pkg/models"
	"github.com/kiyuu10/go-bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookList(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	books = models.GetAllBook()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var (
		book   = &models.Book{}
		params = mux.Vars(r)
		id     = params["id"]
	)
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}

	book, _ = models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book   = models.Book{}
		params = mux.Vars(r)
		id     = params["id"]
	)
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	utils.ParseBody(r, book)
	bookUpdate, _ := models.UpdateBookById(bookId, book)
	res, _ := json.Marshal(bookUpdate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		id     = params["id"]
	)
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	models.DeleteBook(bookId)
	w.WriteHeader(http.StatusOK)
}
