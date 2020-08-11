package module1

// Budget stores budget information
type Budget struct {
	max float32
	Items Item[]
}
// Item stores item information
type Item struct {
	Description string
	Price float32
}