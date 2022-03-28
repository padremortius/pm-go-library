package database

import (
	"pm-go-library/models/authorsModel"
	"pm-go-library/models/genresModel"
	"pm-go-library/models/publishersModel"
	"pm-go-library/models/seriesModel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	My_db gorm.DB
)

func init() {
	My_db = InitDBConn()
}

func InitDBConn() gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/metadata.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return *db
}

func GetRecCountInTable(aTableName string) int64 {
	var count int64
	My_db.Table(aTableName).Count(&count)
	return count
}

func FillSBCListByQuery() seriesModel.SBCList {
	var data seriesModel.SBCList
	sqlQuery1 := "select s.name, count(s.id) as CountBooks from series s inner join books_series_link bsl ON bsl.series = s.id group by s.id order by s.name"
	My_db.Raw(sqlQuery1).Scan(&data)
	return data
}

func FillGBCListByQuery() genresModel.GBCList {
	var data genresModel.GBCList
	sqlQuery1 := "select t.name, count(*) as CountBooks from tags t inner join books_tags_link btl ON btl.tag = t.id group by t.id order by t.name"
	My_db.Raw(sqlQuery1).Scan(&data)
	return data
}

func FillFirstCharOfAuthorsWithCount() authorsModel.AuthorLastNameFCList {
	sqlQuery1 := "select SUBSTRING(sort, 1, 1) as FirstChar, count(*) as AuthorsCount from authors a group by FirstChar"
	var data authorsModel.AuthorLastNameFCList
	My_db.Raw(sqlQuery1).Scan(&data)
	return data
}

func FillFirstCharOfPublishersWithCount() publishersModel.PFCList {
	sqlQuery1 := "select SUBSTRING(name, 1, 1) as FirstChar, count(*) as PublishersCount from publishers a group by FirstChar"
	var data publishersModel.PFCList
	My_db.Raw(sqlQuery1).Scan(&data)
	return data
}
