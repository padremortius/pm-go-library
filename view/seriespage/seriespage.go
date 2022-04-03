package seriespage

import (
	"pm-go-library/common"
	"pm-go-library/database"

	"github.com/gin-gonic/gin"
)

var ()

func ShowSeriesPage(c *gin.Context) {
	s := database.FillSBCListByQuery()
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Серии",
			"payload": s,
		},
		"series.html",
	)
}
