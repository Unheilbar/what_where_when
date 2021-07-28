package service

import (
	"github.com/unheilbar/what_where_when"
	"github.com/unheilbar/what_where_when/pkg/repository"
)

type RoomManager interface {
	CreateRoom(title string, host uint, questionList []what_where_when.Question) error
	RemoveRoom(id uint) error
	GetRoomById(id uint) what_where_when.Room
	GetAllRooms() []what_where_when.Room
}

type Room interface {
	AddPlayer(player what_where_when.Player, roomId uint) error
	RemovePlayer(playerId uint, roomId uint) error
	Guess(roomId, playerId uint, guess string) (bool, error)
	NextQuestion(roomId uint) bool
	AddPointsToPlayer(roomId, playerId, points uint) error
	GetAllPlayers(roomId uint) []what_where_when.Player
}

type Questions interface {
	GetQuestionsList() ([]what_where_when.Question, error)
}

type Services struct {
	RoomManager
	Room
	Questions
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		RoomManager: CreateRoomManagerService(),
		Room:        CreateRoomService(),
		Questions:   CreateQuestionsService(),
	}
}
