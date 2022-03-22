package authorsModel

type (
	AuthorLastNameFCList struct {
		Id        int    `json:"id"`
		FirstChar string `json:"firstchar"`
		Count     int    `json:"count"`
	}
)
