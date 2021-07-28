package repository

import (
	"encoding/json"
	"fmt"
	"time"

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

func (rm *RoomManagerRedis) CreateRoom(title string, host string, questionList []what_where_when.Question) error {
	ex, err := rm.rc.Exists(ctx, title).Result()

	if err != nil {
		return fmt.Errorf("Cant check room existance")
	}

	if ex == 1 {
		return fmt.Errorf("room with title %v exists already", title)
	}

	room := what_where_when.Room{
		Title:        title,
		QuestionList: questionList,
		Host:         host,
		Status:       0,
		PlayerList:   make([]what_where_when.Player, 5),
		Created:      time.Now(),
	}

	js, _ := json.Marshal(room)

	s, err := rm.rc.Set(ctx, title, js, 10*time.Hour).Result()

	if err != nil {
		fmt.Println(s)
		return fmt.Errorf("err occured when setted room %v", err)
	}

	return nil
}

func (rm *RoomManagerRedis) RemoveRoom(title string) error {
	ex, err := rm.rc.Exists(ctx, title).Result()

	if err != nil {
		return fmt.Errorf("Cant check room existance")
	}

	if ex == 0 {
		return fmt.Errorf("room with title %v doesnt exists", title)
	}

	res, err := rm.rc.Get(ctx, title).Result()

	fmt.Println(res)

	res2, err := rm.rc.Del(ctx, title).Result()

	fmt.Printf("result %v", res2)

	return nil
}

func (rm *RoomManagerRedis) GetRoomByTitle(title string) what_where_when.Room {
	return what_where_when.Room{}
}

func (rm *RoomManagerRedis) GetAllRooms() []what_where_when.Room {
	return make([]what_where_when.Room, 0)
}
