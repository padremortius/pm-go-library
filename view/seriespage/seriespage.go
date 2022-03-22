package seriespage

import (
	"pm-go-library/common"
	"pm-go-library/models/seriesModel"

	"github.com/gin-gonic/gin"
)

var seriesList = []seriesModel.SeriesList{
	{FirstChar: "Хроники Амбера", Count: 25},
	{FirstChar: "O'Reily", Count: 10},
	{FirstChar: "ЖЗЛ", Count: 2},
}

func ShowSeriesPage(c *gin.Context) {
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Серии",
			"payload": seriesList,
		}, "series1.html",
	)
}
