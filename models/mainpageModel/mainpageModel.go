package mainpageModel

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Descr string `json:"dedcr"`
	Count int    `json:"count"`
}
