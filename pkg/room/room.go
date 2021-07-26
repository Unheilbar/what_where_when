package room_pool

type Room struct {
	Title       string
	Host        string
	Players     []Player
	QuetionList []Question
	Status      int
}
