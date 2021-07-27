package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/what_where_when/pkg/service"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", h.home)

	router.POST("/create_room", h.createRoom)

	router.GET("/ws", h.ws)

	router.GET("/room/:room", h.joinRoom)

	return router
}
