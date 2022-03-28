package authorsModel

type (
	AuthorLastNameFCItem struct {
		FirstChar    string `json:"firstchar"`
		AuthorsCount int64  `json:"AuthorsCount"`
	}

	AuthorLastNameFCList []AuthorLastNameFCItem
)
