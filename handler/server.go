package handler

import (
	"final-project/config"
	"final-project/manager"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	s.srv.Use(middleware.LoggerMiddleware())

	// handler

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
