package authorspage

import (
	"pm-go-library/common"
	"pm-go-library/database"

	"github.com/gin-gonic/gin"
)

func ShowFFCLPage(c *gin.Context) {
	s := database.FillFirstCharOfAuthorsWithCount()
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Авторы",
			"payload": s,
		}, "authors.html",
	)
}
