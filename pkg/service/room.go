package service

import "github.com/unheilbar/what_where_when"

type RoomService struct {
}

func CreateRoomService() *RoomService {
	return &RoomService{}
}

func (r *RoomService) AddPlayer(player what_where_when.Player, roomId uint) error {
	return nil
}

func (r *RoomService) RemovePlayer(roomId uint, playerId uint) error {
	return nil
}

func (r *RoomService) Guess(roomId, playerId uint, guess string) (bool, error) {
	return true, nil
}

func (r *RoomService) NextQuestion(roomId uint) bool {
	return true
}

func (r *RoomService) AddPointsToPlayer(roomId, playerId, points uint) error {
	return nil
}

func (r *RoomService) GetAllPlayers(roomId uint) []what_where_when.Player {
	return make([]what_where_when.Player, 0)
}
