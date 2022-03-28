package routes

import (
	"pm-go-library/view/aboutpage"
	"pm-go-library/view/authorspage"
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
	Router.GET("/authors1", authorspage.ShowFFCLPage)
	Router.GET("/genres", genrespage.ShowGenresPage)
	Router.GET("/series1", seriespage.ShowSeriesPage)
	Router.GET("/publishers", publisherspage.ShowPFCLPage)
}
