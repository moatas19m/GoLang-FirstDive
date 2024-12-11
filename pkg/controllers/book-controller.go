package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moatas19m/GoLang-FirstDive/pkg/models"
	"github.com/moatas19m/GoLang-FirstDive/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(newBooks)
	if er != nil {
		return
	}
}
