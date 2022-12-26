package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var booksData = []book{
	{ID: 1, Title: "Book1", Price: 56.99},
	{ID: 2, Title: "Book2", Price: 17.99},
	{ID: 3, Title: "Book3", Price: 39.99},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, booksData)
}

func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	booksData = append(booksData, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	validId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Not valid id")
	}

	// Loop over the list of books, looking for
	// an book whose ID value matches the parameter.
	for _, a := range booksData {
		if a.ID == validId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func updateBookByID(c *gin.Context) {
	id := c.Param("id")

	validId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Not valid id")
	}

	var data book
	if err := c.BindJSON(&data); err != nil {
		fmt.Println("Not valid data")
	}

	for i, a := range booksData {
		if a.ID == validId {
			updatedBook := book{validId, data.Title, data.Price}
			booksData[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func RemoveIndex(s []book, index int) []book {
    return append(s[:index], s[index+1:]...)
}

func deleteBookByID(c *gin.Context) {
	id := c.Param("id")

	validId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Not valid id")
	}

	for i, a := range booksData {
		if a.ID == validId {
			booksData = RemoveIndex(booksData, i)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "check",
		})
	})
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.PATCH("books/:id", updateBookByID)
	router.DELETE("/books/:id", deleteBookByID)
	router.Run("localhost:4000") // listen and serve on 0.0.0.0:4000 (for windows "localhost:4000")
}
