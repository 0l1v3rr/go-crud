package data

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Author      string `json:"author"`
}

var Books []Book

func CreateBooks() {
	Books = append(Books, Book{ID: "1", Title: "Book 1", Description: "Book 1 description", Category: "Category 1", Author: "John"})
	Books = append(Books, Book{ID: "2", Title: "Book 2", Description: "Book 2 description", Category: "Category 1", Author: "Jonathan"})
	Books = append(Books, Book{ID: "3", Title: "Book 3", Description: "Book 3 description", Category: "Category 2", Author: "Mike"})
	Books = append(Books, Book{ID: "4", Title: "Book 4", Description: "Book 4 description", Category: "Category 2", Author: "Steve"})
}
