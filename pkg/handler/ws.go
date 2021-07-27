package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ws(c *gin.Context) {
	fmt.Println(c.GetHeader("Cookie"))
	clientToken := strings.Split(c.GetHeader("Cookie"), "=")[1]
	if clientToken == "" {
		c.Status(403)
		return
	}
	//hub.ServeWs(h.services.Hub, c.Writer, c.Request, clientToken)
}
