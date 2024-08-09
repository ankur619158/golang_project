package server

import (
	"fmt"
	"golang_project/config"
	"net/http"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg         config.Config
	router      *gin.Engine
	userHandler userHandler
}

func NewServer(cfg *config.Config, userHandler userHandler) Server {
	router, err := SetupRouter(*cfg)

	if err != nil {
		panic(err)
	}
	return Server{
		cfg:         *cfg,
		router:      router,
		userHandler: userHandler,
	}
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

func (s *Server) SetupInternalRoutes() {
	s.sampleRoutes(s.router.Group("/"))
}

func (s *Server) Run() error {
	
	addr := fmt.Sprintf(":%s", "8000")

	s.SetupInternalRoutes()

	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  14 * time.Minute,
		WriteTimeout: 14 * time.Minute,
	}

	return srv.ListenAndServe()
}

func SetupRouter(cfg config.Config) (*gin.Engine, error) {
	
	r := gin.Default()
	c := cors.DefaultConfig()

	c.AllowOrigins = []string{"*"}
	c.AllowHeaders = []string{"*"}
	c.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT"}
	c.ExposeHeaders = []string{"Content-Disposition"}

	r.Use(cors.New(c))

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         0,
	}))

	return r, nil
}
