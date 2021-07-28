package what_where_when

import "time"

type Room struct {
	Id           int
	Title        string
	QuestionList []Question
	Host         string
	Status       int
	PlayerList   []Player
	Created      time.Time
}
