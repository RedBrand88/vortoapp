package todo

//Item is a single todo item
type Item struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Complete int `json:"complete"`
}