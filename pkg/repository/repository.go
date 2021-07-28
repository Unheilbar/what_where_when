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
	AddPlayer(player what_where_when.Player, roomId uint) error
	RemovePlayer(roomId, playerId uint) error
	Guess(roomId uint, guess string) (bool, error)
	NextQuestion(roomId uint) (bool, what_where_when.Question)
	AddPointsToPlayer(roomId, playerId, points uint)
	GetAllPlayers() []what_where_when.Player
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
