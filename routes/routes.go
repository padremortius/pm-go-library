package routes

import (
	"pm-go-library/view/aboutpage"
	"pm-go-library/view/startpage"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {
	// Handle the index route
	Router.GET("/", startpage.ShowIndexPage)
	Router.GET("/info", aboutpage.ShowAboutPage)
}
