package main

import (
	"pm-go-library/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	routes.Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	routes.Router.LoadHTMLGlob("templates/pages/*")
	routes.Router.Static("/css", "templates/css")

	// Initialize the routes
	routes.InitializeRoutes()

	// Start serving the application
	_ = routes.Router.Run()

}
