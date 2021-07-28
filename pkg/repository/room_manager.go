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

func (rm *RoomManagerRedis) CreateRoom(title string, host uint, questionList []what_where_when.Question) error {
	roomId := rm.rc.Incr(ctx, "room_id")
	return nil
}

func (rm *RoomManagerRedis) RemoveRoom(id uint) error {
	return nil
}

func (rm *RoomManagerRedis) GetRoomById(id uint) what_where_when.Room {
	return what_where_when.Room{}
}

func (rm *RoomManagerRedis) GetAllRooms() []what_where_when.Room {
	return make([]what_where_when.Room, 0)
}
