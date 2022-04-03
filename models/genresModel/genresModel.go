package genresModel

type (
	Genre struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	GenreList []Genre

	//type GenreList with books count
	GenreItem struct {
		GenreId    int    `json:"genreid"`
		GenreName  string `json:"genrename"`
		CountBooks int    `json:"countbooks"`
	}

	GenresList []GenreItem
)
