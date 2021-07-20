package service

import "github.com/unheilbar/what_where_when/pkg/hub"

type Dependencies struct {
	HomeTemplate string
}

type Service struct {
	Hub *hub.Hub
}

func NewService(hub *hub.Hub) *Service {
	return &Service{
		Hub: hub,
	}
}
