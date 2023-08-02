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
	NewCustomerHandler(s.srv, s.usecaseManager.GetCustomerUsecase())
	NewTransactionHandler(s.srv, s.usecaseManager.GetTransactionUsecase())
	NewMerchantHandler(s.srv, s.usecaseManager.GetMerchantUsecase())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()

	repo := manager.NewRepoManager()
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
