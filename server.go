package what_where_when

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
