package booksModel

import "time"

type (
	Book struct {
		Id            int       `json:"id"`
		Title         string    `json:"title"`
		Sort          string    `json:"sort"`
		Timestamp     time.Time `json:"datetime"`
		Pubdate       string    `json:"pubdate"`
		Series_index  int       `json:"series_index"`
		Author_sort   string    `json:"author_sort"`
		Isbn          string    `json:"isbn"`
		Iccn          string    `json:"iccn"`
		Path          string    `json:"path"`
		Flags         int       `json:"flags"`
		Uuid          string    `json:"uuid"`
		Has_cover     bool      `json:"has_cover"`
		Last_modified time.Time `json:"last_modified"`
	}

	Books []Book

	BooksListItem struct {
		Num        int    `json:"num"`
		AuthorName string `json:"author_sort"`
		BookTitle  string `json:"sort"`
		PubDate    string `json:"pubdate"`
	}

	BooksList []BooksListItem
)
