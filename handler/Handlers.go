package ascii

import (
	"net/http"
	"os"
	"strings"
	"text/template"
	"unicode"

	ascii "ascii/functions"
)

var Tp *template.Template

type ErrorPage struct {
	Code         int
	ErrorMessage string
}

func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)

	if err != nil || File.IsDir() {

		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))).ServeHTTP(w, r)
}

func ResultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	if r.Method != http.MethodPost {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	word := r.FormValue("word")
	typee := r.FormValue("typee")

	ctr := 0
	for _, v := range word {
		if v != '\r' {
			ctr++
		}
	}
	var errorMessage string

	if word == "" {
		errorMessage = "Please enter a word."
	} else if typee == "" {
		errorMessage = "Please select a type."
	} else if ctr > 1000 {
		errorMessage = "The word length should not exceed 1000 characters."
	}

	for i := 0; i < len(word); i++ {
		if unicode.IsLetter(rune(word[i])) && (word[i] < 32 || word[i] > 126) {
			errorMessage = "Invalid characters."
			break
		}
	}

	LastResult := ascii.Ascii(word, typee)

	if LastResult == "" {
		errorMessage = " invalid file name !!!!! "
	}

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		Tp.ExecuteTemplate(w, "index.html", errorMessage)
		return
	}

	err := Tp.ExecuteTemplate(w, "result.html", LastResult)
	if err != nil {
		errore := ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Something went wrong on our end. Please try again later.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
	}
}

func FormFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	if r.Method != http.MethodGet {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	err := Tp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		errore := ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Something went wrong on our end. Please try again later.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
	}
}
