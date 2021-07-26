package service

import (
	"fmt"
	"sync"
)

type Room struct {
	Title           string
	Host            string
	Players         map[string]Player
	CurrentQuestion Question
	QuestionList    []Question
	Status          int
}

type RoomManagerService struct {
	mx       sync.Mutex
	RoomPool map[string]Room
}

func CreateRoomManagerService() *RoomManagerService {
	return &RoomManagerService{
		RoomPool: make(map[string]Room, 10),
	}
}

func (rm *RoomManagerService) CreateRoom(title string, host string, questionsList []Question) error {
	rm.mx.Lock()
	defer rm.mx.Unlock()
	_, ok := rm.RoomPool[title]

	if ok {
		return fmt.Errorf("room %v already exists", title)
	}

	rm.RoomPool[title] = Room{
		Title:        title,
		QuestionList: questionsList,
		Host:         host,
		Status:       0, // Waiting for other players
	}
	return nil
}

func (rm *RoomManagerService) DeleteRoom(roomTitle string) error {
	rm.mx.Lock()
	defer rm.mx.Unlock()

	_, ok := rm.RoomPool[roomTitle]

	if ok {
		delete(rm.RoomPool, roomTitle)

		return nil
	}
	return fmt.Errorf("no room %v in room pool", roomTitle)
}

func (rm *RoomManagerService) GetAllRooms() []string {
	res := make([]string, len(rm.RoomPool))

	rm.mx.Lock()
	for title := range rm.RoomPool {
		res = append(res, title)
	}
	rm.mx.Unlock()
	return res
}

func (rm *RoomManagerService) Guess(title string, player string, guess string) (bool, error) {
	rm.mx.Lock()
	defer rm.mx.Unlock()
	_, ok := rm.RoomPool[title]
	if !ok {
		return false, fmt.Errorf("no room %v in room pool", title)
	}

	_, ok = rm.RoomPool[title].Players[player]
	if !ok {
		return false, fmt.Errorf("no player %v in room %v", player, title)
	}

	return match(guess, rm.RoomPool[title].CurrentQuestion.Answer), nil
}

func (rm *RoomManagerService) NextQuestion(title string) bool {
	rm.mx.Lock()
	defer rm.mx.Unlock()

	for rt, room := range rm.RoomPool {
		if rt == title {
			room.QuestionList = room.QuestionList[1:]
			room.CurrentQuestion = room.QuestionList[0]
		}
	}

	return len(rm.RoomPool[title].QuestionList) == 1
}

func match(str1, str2 string) bool {
	return str1 == str2
}

type Player struct {
	Name   string
	Points int
	Host   bool
}

type Question struct {
	Question string
	Category string
	Options  map[string]bool
	Answer   string
}
