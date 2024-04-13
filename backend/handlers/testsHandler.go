package handlers

import (
	"html/template"
	"net/http"
)

func TestsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
