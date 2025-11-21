package structembedding

import "encoding/json"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// Author represents information about the book's author
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Book represents information about a book
type Book struct {
	Title  string `json:"title"`
	// Embedding: itt nincs mezőnév, csak a típus neve. 
	// A json taget viszont kitesszük, hogy tudja a parser, honnan szedje az adatot.
	Author `json:"author"` 
	Pages  int    `json:"pages"`
	ISBN   string `json:"ISBN"`
}

// Article represents information about a article
type Article struct {
	Title   string `json:"title"`
	Author  `json:"author"`
	Journal string `json:"journal"`
	Year    int    `json:"year"`
}

// ParseBook parses the given JSON data into a Book struct
func ParseBook(jsonData []byte) (Book, error) {
	var b Book
	// A json.Unmarshal pointert vár a célstruktúrára (&b)
	err := json.Unmarshal(jsonData, &b)
	return b, err
}

// ParseArticle parses the given JSON data into a Article struct
func ParseArticle(jsonData []byte) (Article, error) {
	var a Article
	err := json.Unmarshal(jsonData, &a)
	return a, err
}
