package publishersModel

type (
	Publisher struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Sort string `json:"sort,omitemty"`
	}

	PublisherList []Publisher

	PublisherItem struct {
		PublisherId     int    `json:"publisherid"`
		PublisherName   string `json:"publishername"`
		PublishersCount int64  `json:"publisherscount"`
	}

	PublishersList []PublisherItem
)
