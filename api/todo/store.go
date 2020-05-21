package todo

//Store is an interface for any obj that will handle requests
type Store interface {
	Get() []Item
	Set(Item) bool
	Update(Item) bool
}