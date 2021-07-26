package service

import "github.com/unheilbar/what_where_when/pkg/hub"

type Dependencies struct {
	HomeTemplate string
}

type RoomManager interface {
	CreateRoom(title string, host string, questionsList []Question) error
	DeleteRoom(title string) error
	GetAllRooms() []string
}

type Service struct {
	RoomManager
	Hub *hub.Hub
}

func NewService(hub *hub.Hub) *Service {
	return &Service{
		Hub:         hub,
		RoomManager: CreateRoomManagerService(),
	}
}
