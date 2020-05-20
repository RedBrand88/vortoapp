package todo

//Item is a single todo item
type Item struct {
	Text string `json:"text"`
	Complete int `json:"complete"`
}