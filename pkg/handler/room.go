package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createRoom(c *gin.Context) {
	title, ok := c.GetPostForm("room-title")

	if !ok {
		c.Redirect(http.StatusPermanentRedirect, "/")
		c.Abort()
	}

	roomUrl := "/room/" + title
	fmt.Println(roomUrl)
	c.Redirect(http.StatusMovedPermanently, roomUrl)
	c.Abort()
}

func (h *Handler) joinRoom(c *gin.Context) {
	roomTitle := c.Param("room")
	fmt.Println(roomTitle)
	c.JSON(200, roomTitle)
}
