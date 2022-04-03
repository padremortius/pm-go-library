package routes

import (
	"pm-go-library/view/aboutpage"
	"pm-go-library/view/authorspage"
	"pm-go-library/view/bookspage"
	"pm-go-library/view/genrespage"
	"pm-go-library/view/publisherspage"
	"pm-go-library/view/seriespage"
	"pm-go-library/view/startpage"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {
	// Handle the index route
	Router.GET("/", startpage.ShowIndexPage)
	Router.GET("/info", aboutpage.ShowAboutPage)
	Router.GET("/authors", authorspage.ShowFFCLPage)
	Router.GET("/genres", genrespage.ShowGenresPage)
	Router.GET("/series", seriespage.ShowSeriesPage)
	Router.GET("/publishers", publisherspage.ShowPFCLPage)
	Router.GET("/booksByFCL", bookspage.ShowBooksByFirstCharOfLastname)
	Router.GET("/booksBySerieId", bookspage.ShowBooksBySerieId)
	Router.GET("/booksByPublisherId", bookspage.ShowBooksByPublisherId)
	Router.GET("/booksByGenreId", bookspage.ShowBooksByGenreId)
}
