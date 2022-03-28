package publisherspage

import (
	"pm-go-library/common"
	"pm-go-library/database"

	"github.com/gin-gonic/gin"
)

func ShowPFCLPage(c *gin.Context) {
	s := database.FillFirstCharOfPublishersWithCount()
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Издательства",
			"payload": s,
		}, "publishers.html",
	)
}
