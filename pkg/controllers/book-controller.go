package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/moatas19m/GoLang-FirstDive/pkg/models"
	"github.com/moatas19m/GoLang-FirstDive/pkg/utils"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Error creating book", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(createdBook)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(newBooks)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(book)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book models.Book
	err = utils.ParseJSONBody(r, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook, err := models.UpdateBookById(id, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(updatedBook)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := models.DeleteBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(book)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}
