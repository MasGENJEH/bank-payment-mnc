package delivery

import (
	"database/sql"
	"fmt"
	"test-mnc/config"
	"test-mnc/delivery/controller"
	"test-mnc/delivery/middleware"
	"test-mnc/repository"
	"test-mnc/shared/service"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	customerUC usecase.CustomersUseCase
	transactionUC usecase.TransactionsUsecase
	authUsc        usecase.AuthUseCase
	engine         *gin.Engine
	jwtService     service.JwtService
	host           string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUsc, rg).Route()
	controller.NewTransactionsController(s.transactionUC, rg, authMiddleware).Route()

}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// Inject DB ke -> repository
	customerRepo := repository.NewCustomerRepository(db)
	transactionRepo := repository.NewTransactionsRepository(db)

	// Inject REPO ke -> useCase
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	transactionUC := usecase.NewTransactionsUsecase(transactionRepo)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUc := usecase.NewAuthUseCase(customerUC, jwtService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		customerUC: customerUC,
		transactionUC: transactionUC,
		authUsc:        authUc,
		engine:         engine,
		jwtService:     jwtService,
		host:           host,
	}
}