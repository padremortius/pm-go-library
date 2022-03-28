package genresModel

type (
	Genre struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	GenreList []Genre

	//type GenreList with books count
	GBCItem struct {
		Name       string `json:"name"`
		CountBooks int    `json:"countbooks"`
	}

	GBCList []GBCItem
)
