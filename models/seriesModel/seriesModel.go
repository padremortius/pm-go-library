package seriesModel

type (
	Serie struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Sort       string `json:"sort"`
		CountBooks int    `json:"countBooks"`
	}

	SeriesList []Serie

	//type SeriesList with books count
	SBCItem struct {
		Name       string `json:"name"`
		CountBooks int    `json:"countbooks"`
	}

	SBCList []SBCItem
)
