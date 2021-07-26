package room_pool

import (
	"fmt"
	"sync"
)

type RoomManager struct {
	mx       sync.Mutex
	RoomPool map[string]Room
}

func CreateRoomPool() *RoomManager {
	return &RoomManager{
		RoomPool: make(map[string]Room, 10),
	}
}

func (rp *RoomManager) CreateRoom(roomTitle string, questionsList []Question) error {
	rp.mx.Lock()
	defer rp.mx.Unlock()
	_, ok := rp.RoomPool[roomTitle]

	if ok {
		return fmt.Errorf("room %v already exists", roomTitle)
	}

	rp.RoomPool[roomTitle] = Room{
		Title:       roomTitle,
		QuetionList: questionsList,
		Status:      0, // Waiting for other players
	}
	return nil

}

func (rp *RoomManager) RemoveRoom(roomTitle string) error {
	rp.mx.Lock()
	defer rp.mx.Unlock()

	_, ok := rp.RoomPool[roomTitle]

	if ok {
		delete(rp.RoomPool, roomTitle)

		return nil
	}
	return fmt.Errorf("no room %v in room pool", roomTitle)
}

func (rp *RoomManager) GetAllRoomPool() []string {
	res := make([]string, len(rp.RoomPool))

	rp.mx.Lock()
	for title := range rp.RoomPool {
		res = append(res, title)
	}
	rp.mx.Unlock()
	return res
}
