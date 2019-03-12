package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func init() {
	// Mock data - @todo - implement DB
	books = append(books, Book{
		ID:     "1",
		ISBN:   "12345",
		Title:  "Book one",
		Author: &Author{Firstname: "John", Lastname: "Doe"},
	})
	books = append(books, Book{
		ID:     "2",
		ISBN:   "67889",
		Title:  "Book two",
		Author: &Author{Firstname: "Bob", Lastname: "Ross"},
	})

}

func main() {
	// init mux router
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// http.Handle("/", r)
	// server := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "127.0.0.1:8080",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Printf("Listening on: %v\n", server.Addr)
	// log.Fatal(server.ListenAndServe())

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
