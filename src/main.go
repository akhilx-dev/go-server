package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Book struct
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

var books []Book

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	for _, item := range books {
		if item.ID == parameters["id"] {
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	writer.Write([]byte("No book found!\n"))
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Deleted book successfully."))
	writer.Write([]byte("Id not found"))

}

func main() {
	books = append(books, Book{"100", "600001", "Rich Dad Poor Dad", Author{
		"Robert",
		"Kiosaki",
	}})
	books = append(books, Book{"201", "700002", "As You Like It", Author{
		"William",
		"Shakespeare",
	}})
	books = append(books, Book{"303", "800003", "An American Tragedy", Author{
		"Theodore", "Dreiser",
	}})

	fmt.Println("Hi there", books)
	router := mux.NewRouter()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books?id={id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
