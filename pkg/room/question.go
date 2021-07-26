package room_pool

type Question struct {
	Question string
	Category string
	Options  map[string]bool
	Answer   string
}
