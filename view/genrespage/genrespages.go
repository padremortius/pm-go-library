package genrespage

import (
	"pm-go-library/common"
	"pm-go-library/database"

	"github.com/gin-gonic/gin"
)

func ShowGenresPage(c *gin.Context) {
	s := database.FillGBCListByQuery()
	common.Render(
		c,
		gin.H{
			"title":   "PM-Go-Library: Жанры",
			"payload": s,
		}, "genres.html",
	)
}
