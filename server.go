package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("Templates/*.html")
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	apiKey := "PfJlPbMJNn7XYdA5FgOOPl22fthxOb9mwT17Nho1"
	photo, err := FetchNasaPhoto(apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch NASA photo", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", photo)
}
