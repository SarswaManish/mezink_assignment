package server

import (
	"awesomeProject/configs"
	"awesomeProject/internal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerResources struct {
	DB            *gorm.DB
	Cors          configs.Cors
	RecordHandler internal.RecordHandler
	InternalAuth  struct {
		Username string
		Password string
	}
}

type Server struct {
	ServerResources
}

type ErrInternalServer struct {
	Err error
}

func (e ErrInternalServer) Error() string {
	return e.Err.Error()
}

func NewServer(res ServerResources) *Server {
	return &Server{
		ServerResources: res,
	}
}

func (s *Server) KnockKnock(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"success": true,
		"message": "mez_ink",
	})
}
