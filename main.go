package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book represents data about a record Book.
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Books slice to seed record Book data.
var books = []book{
	{ID: "1", Title: "Salems lot", Author: "Steven King", Price: 21.99},
	{ID: "2", Title: "The Priory of the Orange Tree", Author: "Samantha Shannon", Price: 29.99},
	{ID: "3", Title: "1984", Author: "George Orwell", Price: 9.99},
	{ID: "4", Title: "The Fellowship of the Ring ", Author: "J.R.R. Tolkien", Price: 39.99},
	{ID: "5", Title: "The Two Towers", Author: "J.R.R. Tolkien", Price: 39.99},
	{ID: "6", Title: "The Return of the King", Author: "J.R.R. Tolkien", Price: 39.99},
}

// getBooks responds with the list of all Books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds an Book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new Book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the Book whose ID value matches the id
// parameter sent by the client, then returns that Book as a response.
func getBookbyID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of Books, looking for
	// an Book whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookbyID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}
