package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Book struct represents a book with basic information.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Mock database to store books
var books []Book

// Handler to get all books (API)
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Handler to get a single book by ID (API)
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to create a new book (API)
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Handler to update a book by ID (API)
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = params["id"]
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to delete a book by ID (API)
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// CLI Functions

func showMenu() {
	fmt.Println("\n--- Book Management CLI ---")
	fmt.Println("1. View all books")
	fmt.Println("2. View a book by ID")
	fmt.Println("3. Add a new book")
	fmt.Println("4. Update a book by ID")
	fmt.Println("5. Delete a book by ID")
	fmt.Println("6. Exit CLI")
	fmt.Print("Select an option: ")
}

func viewAllBooksCLI() {
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}
	fmt.Println("\nList of Books:")
	for _, book := range books {
		fmt.Printf("ID: %s | Title: %s | Author: %s | Year: %d\n", book.ID, book.Title, book.Author, book.Year)
	}
}

func viewBookByIDCLI() {
	id := input("Enter book ID: ")
	for _, book := range books {
		if book.ID == id {
			fmt.Printf("ID: %s | Title: %s | Author: %s | Year: %d\n", book.ID, book.Title, book.Author, book.Year)
			return
		}
	}
	fmt.Println("Book not found.")
}

func addBookCLI() {
	book := Book{
		ID:     input("Enter book ID: "),
		Title:  input("Enter book title: "),
		Author: input("Enter book author: "),
		Year:   intInput("Enter publication year: "),
	}
	books = append(books, book)
	fmt.Println("Book added successfully.")
}

func updateBookByIDCLI() {
	id := input("Enter book ID to update: ")
	for i, book := range books {
		if book.ID == id {
			books[i].Title = input("Enter new title: ")
			books[i].Author = input("Enter new author: ")
			books[i].Year = intInput("Enter new publication year: ")
			fmt.Println("Book updated successfully.")
			return
		}
	}
	fmt.Println("Book not found.")
}

func deleteBookByIDCLI() {
	id := input("Enter book ID to delete: ")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Println("Book deleted successfully.")
			return
		}
	}
	fmt.Println("Book not found.")
}

// Helper functions for CLI
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func intInput(prompt string) int {
	var num int
	fmt.Print(prompt)
	fmt.Scanf("%d", &num)
	return num
}

// Start CLI
func startCLI() {
	for {
		showMenu()
		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			viewAllBooksCLI()
		case 2:
			viewBookByIDCLI()
		case 3:
			addBookCLI()
		case 4:
			updateBookByIDCLI()
		case 5:
			deleteBookByIDCLI()
		case 6:
			fmt.Println("Exiting CLI.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func main() {
	// Sample data for testing
	books = append(books, Book{ID: "1", Title: "Go Programming", Author: "John Doe", Year: 2020})
	books = append(books, Book{ID: "2", Title: "Advanced Go", Author: "Jane Doe", Year: 2021})

	// Start API server in a separate goroutine
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/books", getBooks).Methods("GET")
		r.HandleFunc("/books/{id}", getBook).Methods("GET")
		r.HandleFunc("/books", createBook).Methods("POST")
		r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
		r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
		fmt.Println("Starting API server on :8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	// Start CLI
	startCLI()
}
