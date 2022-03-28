package startpage

import (
	"pm-go-library/common"
	"pm-go-library/database"
	"pm-go-library/models/mainpageModel"

	"github.com/gin-gonic/gin"
)

func FillItemStartPage() mainpageModel.StartPageList {

	var itemList mainpageModel.StartPageList

	itemList = append(itemList, mainpageModel.Item{
		ID: "authors1", Title: "Авторы", Descr: "Список авторов в алфавитном порядке", Count: database.GetRecCountInTable("authors")})
	itemList = append(itemList, mainpageModel.Item{
		ID: "series1", Title: "Серии", Descr: "Список серий в алфавитном порядке", Count: database.GetRecCountInTable("series")})
	itemList = append(itemList, mainpageModel.Item{
		ID: "publishHouses", Title: "Издательства", Descr: "Список издательств в алфавитном порядке", Count: database.GetRecCountInTable("publishers")})
	itemList = append(itemList, mainpageModel.Item{
		ID: "genres", Title: "Жанры", Descr: "Список жанров в алфавитном порядке", Count: database.GetRecCountInTable("tags")})

	return itemList
}

func ShowIndexPage(c *gin.Context) {
	s := FillItemStartPage()
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library",
			"payload": s,
		}, "index.html",
	)
}
