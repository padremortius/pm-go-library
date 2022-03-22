package startpage

import (
	"pm-go-library/common"
	"pm-go-library/models/mainpageModel"

	"github.com/gin-gonic/gin"
)

var ItemList = []mainpageModel.Item{
	{ID: "authors1", Title: "Авторы", Descr: "Список авторов в алфавитном порядке", Count: 10},
	{ID: "series1", Title: "Серии", Descr: "Список серий в алфавитном порядке", Count: 12},
	{ID: "publishHouses", Title: "Издательства", Descr: "Список издательств в алфавитном порядке", Count: 3},
	{ID: "genres", Title: "Жанры", Descr: "Список жанров в алфавитном порядке", Count: 5},
}

func ShowIndexPage(c *gin.Context) {
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library",
			"payload": ItemList,
		}, "index.html",
	)
}
