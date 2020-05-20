package todo

//Item is a single todo item
type Item struct {
	Text string `json:"string"`
	Complete int `json:"complete"`
}