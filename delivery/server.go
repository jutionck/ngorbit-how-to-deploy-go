package delivery

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/blog-apps/config"
	"enigmacamp.com/blog-apps/delivery/controller"
	"enigmacamp.com/blog-apps/repository"
	"enigmacamp.com/blog-apps/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	blogUC usecase.BlogUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	controller.NewBlogController(s.blogUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}
	blogRepo := repository.NewBlogRepository(db)
	blogUC := usecase.NewBlogUseCase(blogRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		blogUC: blogUC,
		engine: engine,
		host:   host,
	}
}
