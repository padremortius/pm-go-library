package publishersModel

type (
	Publisher struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Sort string `json:"sort,omitemty"`
	}

	PublisherList []Publisher

	PFCItem struct {
		FirstChar       string `json:"firstchar"`
		PublishersCount int64  `json:"publisherscount"`
	}

	PFCList []PFCItem
)
