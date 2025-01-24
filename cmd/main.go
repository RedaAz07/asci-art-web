package main

import (
	"fmt"
	"net/http"
	"text/template"

	handler "ascii/handler"
)

func main() {
	var err error
	handler.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Printf("Error parsing templates: %v\n", err)
		return
	}

	// Register handlers
	http.HandleFunc("/styles/", handler.StyleFunc)
	http.HandleFunc("/ascii-art", handler.ResultFunc)
	http.HandleFunc("/", handler.FormFunc)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
