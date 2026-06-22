package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Input  string
	Banner string
}

type ErrorData struct {
	Code    int
	Message string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("Server successfully started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Fallback baseline if even the error template is missing
		http.Error(w, fmt.Sprintf("%d - %s", status, message), status)
		return
	}
	tmpl.Execute(w, ErrorData{Code: status, Message: message})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Guardrail: Strict routing validation. If path isn't precisely "/", return 404.
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error: Missing UI templates")
		return
	}
	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Parse custom post forms safely
	if err := r.ParseForm(); err != nil {
		renderError(w, http.StatusBadRequest, "Bad Request: Unable to parse form data")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Input Validation
	if text == "" || banner == "" {
		renderError(w, http.StatusBadRequest, "Bad Request: Submissions cannot contain empty elements")
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		renderError(w, http.StatusBadRequest, "Bad Request: Unknown banner style selected")
		return
	}

	// Execute core processing 
	output, statusCode := ascii.GenerateASCII(text, banner)
	if statusCode != 200 {
		// DEBUG PRINT: This will print the internal issue to your terminal console
		fmt.Printf("[SERVER DEBUG] Text input: %q | Banner requested: %q | Internal Status Code: %d\n", text, banner, statusCode)

		switch statusCode {
		case 400:
			renderError(w, http.StatusBadRequest, "Bad Request: Text contains non-ASCII characters")
		case 404:
			renderError(w, http.StatusNotFound, "Internal Error: Banner file not found on disk")
		default:
			renderError(w, http.StatusInternalServerError, "Internal Server Error during layout generation")
		}
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error: Missing UI templates")
		return
	}

	tmpl.Execute(w, PageData{
		Result: output,
		Input:  text,
		Banner: banner,
	})
}