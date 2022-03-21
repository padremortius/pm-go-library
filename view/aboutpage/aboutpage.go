package aboutpage

import (
	"pm-go-library/common"
	"pm-go-library/models/aboutModel"

	"github.com/gin-gonic/gin"
)

var Info = aboutModel.AboutInfo{Version: "0.0.1", Author: "padremortius"}

func ShowAboutPage(c *gin.Context) {
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library",
			"version": Info.Version,
			"author":  Info.Author,
		}, "about.html",
	)
}
