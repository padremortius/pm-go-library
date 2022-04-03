package database

import (
	"fmt"
	"pm-go-library/models/authorsModel"
	"pm-go-library/models/booksModel"
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
	sqlQuery := "select s.id as SerieId, s.name, count(s.id) as CountBooks from series s " +
		"inner join books_series_link bsl ON bsl.series = s.id group by s.id order by s.name"
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillGenreListByQuery() genresModel.GenresList {
	var data genresModel.GenresList
	sqlQuery := "select t.id as GenreId, t.name GenreName, count(*) as CountBooks from tags t " +
		"inner join books_tags_link btl ON btl.tag = t.id group by t.id"
	My_db.Raw(sqlQuery).Order("TagName").Scan(&data)
	return data
}

func FillFirstCharOfAuthorsWithCount() authorsModel.AuthorLastNameFCList {
	sqlQuery := "select SUBSTRING(a.sort, 1, 1) as FirstChar, count(*) as AuthorsCount from authors a " +
		"inner join books_authors_link bal ON bal.author = a.id " +
		"group by FirstChar"
	var data authorsModel.AuthorLastNameFCList
	fmt.Println(sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillPublishersWithCount() publishersModel.PublishersList {
	sqlQuery := "select p.id as PublisherId, p.name as PublisherName, count(*) as PublishersCount from publishers p " +
		"inner join books_publishers_link bpl ON bpl.publisher = p.id " +
		"group by p.id order by p.name"
	var data publishersModel.PublishersList
	fmt.Println(sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillBooksByFirstCharOfLastnameAuthor(items []string) booksModel.BooksList {
	var data booksModel.BooksList
	sqlQuery := "select row_number() over(ORDER BY a.name, b.sort) num, a.name as AuthorName, b.sort as BookTitle, strftime('%d-%m-%Y', b.\"timestamp\") as PubDate " +
		"from books b, books_authors_link bal, authors a " +
		"where bal.book = b.id AND a.id = bal.author AND SUBSTRING(a.sort, 1, 1) = \"" + items[0] + "\""
	fmt.Println(sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillBooksBySerieId(items []string) booksModel.BooksList {
	fmt.Println("FillBooksBySierieId")
	var data booksModel.BooksList
	sqlQuery := "select row_number() over(ORDER BY a.name, b.sort) num, a.name as AuthorName, b.sort as BookTitle, strftime('%d.%m.%Y', b.'timestamp') as PubDate from books b " +
		"inner join books_authors_link bal ON b.id = bal.book " +
		"inner join authors a ON a.id = bal.author " +
		"inner join books_series_link bsl ON bsl.book = b.id " +
		"inner join series s ON s.id = bsl.series " +
		"where s.id = " + items[0]
	fmt.Println("sqlQuery = ", sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillBooksByPublisherId(items []string) booksModel.BooksList {
	fmt.Println("FillBooksByPublisherId")
	var data booksModel.BooksList
	sqlQuery := "select row_number() over(ORDER BY a.name, b.sort) num, a.name as AuthorName, b.sort as BookTitle, strftime('%d.%m.%Y', b.'timestamp') as PubDate from books b " +
		"inner join books_authors_link bal ON b.id = bal.book " +
		"inner join authors a ON a.id = bal.author " +
		"inner join books_publishers_link bpl ON bpl.book = b.id " +
		"inner join publishers p ON p.id = bpl.publisher " +
		"where p.id = " + items[0]
	fmt.Println(sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}

func FillBooksByGenreId(items []string) booksModel.BooksList {
	fmt.Println("FillBooksByGenreId")
	var data booksModel.BooksList
	sqlQuery := "select row_number() over(ORDER BY a.name, b.sort) num, a.name as AuthorName, b.sort as BookTitle, strftime('%d.%m.%Y', b.'timestamp') as PubDate from books b " +
		"inner join books_authors_link bal ON b.id = bal.book " +
		"inner join authors a ON a.id = bal.author " +
		"inner join books_tags_link btl ON btl.book = b.id " +
		"inner join tags t ON t.id = btl.tag " +
		"where t.id = " + items[0]
	fmt.Println(sqlQuery)
	My_db.Raw(sqlQuery).Scan(&data)
	return data
}
