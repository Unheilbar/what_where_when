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
	AddPlayer(nickname string) error
	RemovePlayer(id uint) error
	Guess(playerId uint, guess string) (bool, error)
	NextQuestion() bool
	AddPointsToPlayer(id, points uint)
	GetTopPlayers() []what_where_when.Player
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
