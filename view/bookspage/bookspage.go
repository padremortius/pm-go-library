package bookspage

import (
	"fmt"
	"pm-go-library/common"
	"pm-go-library/database"
	"pm-go-library/models/booksModel"

	"github.com/gin-gonic/gin"
)

func ShowBooksByFirstCharOfLastname(c *gin.Context) {
	fmt.Println("ShowBooksByFirstCharOfLastname")
	var s booksModel.BooksList
	queryParams := c.Request.URL.Query()
	s = database.FillBooksByFirstCharOfLastnameAuthor(queryParams["AutorsFirstLetter"])
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Книги",
			"payload": s,
		}, "books.html",
	)
}

func ShowBooksBySerieId(c *gin.Context) {
	fmt.Println("ShowBooksBySerieId")
	var s booksModel.BooksList
	queryParams := c.Request.URL.Query()
	s = database.FillBooksBySerieId(queryParams["SerieId"])
	fmt.Printf("%+v", s)
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Книги",
			"payload": s,
		},
		"books.html",
	)
}

func ShowBooksByPublisherId(c *gin.Context) {
	fmt.Println("ShowBooksByPublisherId")
	var s booksModel.BooksList
	queryParams := c.Request.URL.Query()
	s = database.FillBooksByPublisherId(queryParams["PublisherId"])
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Книги",
			"payload": s,
		},
		"books.html",
	)
}

func ShowBooksByGenreId(c *gin.Context) {
	fmt.Println("ShowBooksByGenreId")
	var s booksModel.BooksList
	queryParams := c.Request.URL.Query()
	s = database.FillBooksByGenreId(queryParams["GenreId"])
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Книги",
			"payload": s,
		},
		"books.html",
	)
}
