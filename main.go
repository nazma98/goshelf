package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📚 Welcome to GoShelf API!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	fmt.Println("🚀 Server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("❌ Server failed to run", err)
	}
}
