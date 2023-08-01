package handler

import (
	"project/config"
	"project/manager"
	"project/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	// session
	store := cookie.NewStore([]byte("secret"))
	s.srv.Use(sessions.Sessions("session", store))

	s.srv.Use(middleware.LoggerMiddleware())

	// handler
	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())

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
