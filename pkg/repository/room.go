package repository

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/unheilbar/what_where_when"
)

type RoomRedis struct {
	rc *redis.Client
}

func CreateRoomRedis(rc *redis.Client) *RoomRedis {
	return &RoomRedis{
		rc,
	}
}

func (r *RoomRedis) AddPlayer(player what_where_when.Player, roomId uint) error {
	return nil
}

func (r *RoomRedis) RemovePlayer(roomId, playerId uint) error {
	return nil
}

func (r *RoomRedis) Guess(roomId uint, guess string) (bool, error) {
	return false, nil
}

func (r *RoomRedis) NextQuestion(roomId uint) (bool, what_where_when.Question) {
	return true, what_where_when.Question{}
}

func (r *RoomRedis) AddPointsToPlayer(roomId, playerId, points uint) {
	return
}

func (r *RoomRedis) GetAllPlayers() []what_where_when.Player {
	return make([]what_where_when.Player, 0)
}
