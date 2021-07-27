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

func (r *RoomRedis) AddPlayer(nickname string) error

func (r *RoomRedis) RemovePlayer(id uint) error

func (r *RoomRedis) Guess(playerId uint, guess string) (bool, error)

func (r *RoomRedis) NextQuestion() bool

func (r *RoomRedis) AddPointsToPlayer(id, points uint)

func (r *RoomRedis) GetTopPlayers() []what_where_when.Player
