package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"unicode"

	ascii "ascii/functions"
)

func main() {
	http.HandleFunc("/styles/", styleFunc)

	http.HandleFunc("/ascii-art", ResultFunc)
	http.HandleFunc("/", formFunc)
	http.HandleFunc("/notfound", notFoundFunc)
	http.HandleFunc("/MethodNotAllowed", MethodNotAllowedFunc)
	http.HandleFunc("/internalServer", internalServerFunc)

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func styleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/styles/" || !strings.HasSuffix(r.URL.Path, ".css") {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}
	http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))).ServeHTTP(w, r)
}

func formFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	tp2, _ := template.ParseFiles("template/index.html")
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/MethodNotAllowed", http.StatusFound)
		return
	}

	tp2.Execute(w, nil)
}

func ResultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/MethodNotAllowed", http.StatusFound)
		return
	}

	word := r.FormValue("word")
	typee := r.FormValue("typee")

	var errorMessage string

	if word == "" {
		errorMessage = "Please enter a word."
	} else if typee == "" {
		errorMessage = "Please select a type."
	} else if len(word) > 1000 {
		errorMessage = "The word length should not exceed 1000 characters."
	} else {
		for i := 0; i < len(word); i++ {
			if unicode.IsLetter(rune(word[i])) && (word[i] < 32 || word[i] > 126) {
				errorMessage = "invalid charts"
				break
			}
		}
	}

	if errorMessage != "" {
		tp1, _ := template.ParseFiles("template/index.html")
		w.WriteHeader(http.StatusBadRequest)
		tp1.Execute(w, errorMessage)

		return
	}
	LastResult := ascii.Ascii(word, typee)

	if LastResult == "" {
		http.Redirect(w, r, "/internalServer", http.StatusFound)
		return

	}

	tp2, _ := template.ParseFiles("template/result.html")

	tp2.Execute(w, LastResult)
}

func notFoundFunc(w http.ResponseWriter, r *http.Request) {
	tp, _ := template.ParseFiles("template/notfound.html")
	w.WriteHeader(http.StatusNotFound)
	tp.Execute(w, nil)
}

func MethodNotAllowedFunc(w http.ResponseWriter, r *http.Request) {
	tp, _ := template.ParseFiles("template/MethodNotAllowed.html")
	w.WriteHeader(http.StatusMethodNotAllowed)
	tp.Execute(w, nil)
}

func internalServerFunc(w http.ResponseWriter, r *http.Request) {
	tp, _ := template.ParseFiles("template/internalServer.html")
	w.WriteHeader(http.StatusInternalServerError)
	tp.Execute(w, nil)
}
