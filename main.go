package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

var books = []book{
    {ID: "1", Title: "Book 1", Author: "Author 1", Price: 19.99},
    {ID: "2", Title: "Book 2", Author: "Author 2", Price: 24.99},
    {ID: "3", Title: "Book 3", Author: "Author 3", Price: 29.99},
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, b := range books {
        if b.ID == id {
            c.IndentedJSON(http.StatusOK, b)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
    var newBook book

    if err := c.BindJSON(&newBook); err != nil {
        return 
    }

    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", createBook)

    router.Run("localhost:8080")
}
