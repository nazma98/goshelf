package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📚 Welcome to GoShelf API!")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "PLease give a GET method", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(bookList)
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var bookList []Book

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	fmt.Println("🚀 Server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("❌ Server failed to run", err)
	}
}

func init() {
	book1 := Book{
		ID:     1,
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
	}
	book2 := Book{
		ID:     1,
		Title:  "Introducing Go",
		Author: "Caleb Doxsey",
	}

	bookList = append(bookList, book1)
	bookList = append(bookList, book2)

}
