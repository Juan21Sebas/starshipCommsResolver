package starshipcommsresolver

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.uber.org/dig"
)

type Server struct {
	engine    *gin.Engine
	container *dig.Container
}

func NewServer(engine *gin.Engine, container *dig.Container) *Server {
	return &Server{
		engine:    engine,
		container: container,
	}
}

func CreateServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,DELETE,PUT",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))
	return server
}

func (server *Server) Start() error {
	port := "8080" //
	address := fmt.Sprintf(":%s", port)
	return server.engine.Run(address)
}
