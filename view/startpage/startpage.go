package startpage

import (
	"pm-go-library/common"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Descr string `json:"dedcr"`
	Count int    `json:"count"`
}

var ItemList = []Item{
	{ID: "authors", Title: "Авторы", Descr: "Список авторов в алфавитном порядке"},
	{ID: "series", Title: "Серии", Descr: "Список серий в алфавитном порядке"},
	{ID: "publishHouses", Title: "Издательства", Descr: "Список издательств в алфавитном порядке"},
	{ID: "genres", Title: "Жанры", Descr: "Список жанров в алфавитном порядке"},
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
