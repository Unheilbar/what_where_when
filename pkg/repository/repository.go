package repository

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/unheilbar/what_where_when"
)

type RoomManager interface {
	CreateRoom(title string, host uint, questionList []what_where_when.Question) error
	RemoveRoom(id uint) error
	GetRoomById(id uint) what_where_when.Room
	GetAllRooms() []what_where_when.Room
}

type Room interface {
	AddPlayer(nickname string) error
	RemovePlayer(id uint) error
	Guess(playerId uint, guess string) (bool, error)
	NextQuestion() bool
	AddPointsToPlayer(id, points uint)
	GetTopPlayers() []what_where_when.Player
}

type Repository struct {
	RoomManager
	Room
}

func NewRepository(c *redis.Client) *Repository {
	return &Repository{
		RoomManager: CreateRoomManagerRedis(c),
		Room:        CreateRoomRedis(c),
	}
}
