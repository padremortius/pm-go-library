package authorspage

import (
	"pm-go-library/common"
	"pm-go-library/models/authorsModel"

	"github.com/gin-gonic/gin"
)

var FamilyFirstCharList = []authorsModel.AuthorLastNameFCList{
	{FirstChar: "A", Count: 15},
	{FirstChar: "B", Count: 25},
	{FirstChar: "D", Count: 10},
	{FirstChar: "Ж", Count: 2},
}

func ShowFFCLPage(c *gin.Context) {
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Авторы",
			"payload": FamilyFirstCharList,
		}, "authors1.html",
	)
}
