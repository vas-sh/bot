package handlers

import (
	"github.com/gin-gonic/gin"
)

type srv struct {
	engine *gin.Engine
}

func New() *srv {
	r := gin.Default()
	return &srv{
		engine: r,
	}
}

func (s *srv) Router() *gin.RouterGroup {
	return s.engine.Group("/api/")
}

func (s *srv) Run() error {
	return s.engine.Run()
}
