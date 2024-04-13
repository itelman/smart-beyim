package handlers

import (
	"net/http"
	"smart-beyim/models"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, models.ErrorConstructor(statusCode))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
