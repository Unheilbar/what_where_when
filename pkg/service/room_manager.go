package service

import "github.com/unheilbar/what_where_when"

type RoomManagerService struct {
}

func CreateRoomManagerService() *RoomManagerService {
	return &RoomManagerService{}
}

func (rm *RoomManagerService) CreateRoom(title string, host uint, questionList []what_where_when.Question) error {
	return nil
}

func (rm *RoomManagerService) RemoveRoom(id uint) error {
	return nil
}

func (rm *RoomManagerService) GetRoomById(id uint) what_where_when.Room {
	return what_where_when.Room{}
}

func (rm *RoomManagerService) GetAllRooms() []what_where_when.Room {
	return make([]what_where_when.Room, 0)
}
