# 📚 GoShelf

GoShelf is a simple Go-based REST API to manage a library of books. It demonstrates how to build a basic HTTP server and perform CRUD operations using Go's standard library.

---

## 🚀 Features

-   📖 List all books
-   ➕ Add a new book
-   🛠️ Update a book
-   ❌ Delete a book
-   🧪 In-memory storage (no external DB required)

---

## 🏗️ Project Structure

```
goshelf/
├── go.mod # Go module file
├── main.go # Entry point
├── handlers/ # HTTP route handlers
│ └── book.go
├── models/ # Data models
│ └── book.go
└── storage/ # In-memory storage logic
└── memory.go
```

---

## ⚙️ Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/nazma98/goshelf.git
cd goshelf
```

### 2. Initialize Go Modules

go mod tidy

### 3. Run the Server

```
go run main.go
```

The server runs at: http://localhost:8080

## 📬 Example Endpoints

| Method | Endpoint      | Description    |
| ------ | ------------- | -------------- |
| GET    | `/books`      | List all books |
| POST   | `/books`      | Add a new book |
| PUT    | `/books/{id}` | Update a book  |
| DELETE | `/books/{id}` | Delete a book  |

## 🛠 Built With

-   Go — The Go programming language
-   net/http — Go's built-in HTTP server

## 📄 License

This project is open source and available under the MIT License.
