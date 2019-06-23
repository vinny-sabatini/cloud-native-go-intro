package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author, and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON to handle marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to handle unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// Books an initial collection of books
var Books = []Book{
	Book{Title: "Please read my book", Author: "A desperate person", ISBN: "256720348"},
	Book{Title: "Don't worry, cloud is easy", Author: "A super smart person", ISBN: "672309867"},
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
