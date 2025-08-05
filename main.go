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

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Please give a POST method", 400)
		return
	}

	// r.Body => id, title, author => book instance => booklist append

	var newBook Book

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give a valid json", 400)
	}

	newBook.ID = len(bookList) + 1

	bookList = append(bookList, newBook)

	encoder := json.NewEncoder(w)
	encoder.Encode(newBook)

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

	mux.HandleFunc("/books", getBooks)

	mux.HandleFunc("/create-books", createBook)

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
		ID:     2,
		Title:  "Introducing Go",
		Author: "Caleb Doxsey",
	}

	bookList = append(bookList, book1)
	bookList = append(bookList, book2)

}
