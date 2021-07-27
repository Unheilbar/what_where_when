package repository

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/unheilbar/what_where_when"
)

type RoomManagerRedis struct {
	rc *redis.Client
}

func CreateRoomManagerRedis(rc *redis.Client) *RoomManagerRedis {
	return &RoomManagerRedis{
		rc,
	}
}

func (rm *RoomManagerRedis) CreateRoom(title string, host uint, questionList []what_where_when.Question) error

func (rm *RoomManagerRedis) RemoveRoom(id uint) error

func (rm *RoomManagerRedis) GetRoomById(id uint) what_where_when.Room

func (rm *RoomManagerRedis) GetAllRooms() []what_where_when.Room
