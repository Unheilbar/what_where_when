package service

import "github.com/unheilbar/what_where_when"

type RoomService struct {
}

func CreateRoomService() *RoomService {
	return &RoomService{}
}

func (r *RoomService) AddPlayer(nickname string) error {
	return nil
}

func (r *RoomService) RemovePlayer(id uint) error {
	return nil
}

func (r *RoomService) Guess(playerId uint, guess string) (bool, error) {
	return true, nil
}

func (r *RoomService) NextQuestion() bool {
	return true
}

func (r *RoomService) AddPointsToPlayer(id, points uint) {

}

func (r *RoomService) GetTopPlayers() []what_where_when.Player {
	return make([]what_where_when.Player, 0)
}
