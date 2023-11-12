package starshipcommsresolver

import (
	"fmt"

	"starshipCommsResolver/internal/infrastructure/errors"

	"github.com/gin-gonic/gin"

	"go.uber.org/dig"
)

func (server *Server) TopSecretSplitRoutes(router *gin.Engine) error {
	c, err := getTopSecretSplitHandlerController(server.container)
	if err != nil {
		return err
	}
	routes := router
	{
		routes.POST("/topsecret_split/:satellite_name", c.topsecretSplit)
		routes.GET("/topsecret_split/:satellite_name", c.GetsecretSplit)
	}
	return nil
}

func getTopSecretSplitHandlerController(container *dig.Container) (*topsecretSplitHandler, error) {
	var c *topsecretSplitHandler

	if err := container.Invoke(func(sc *topsecretSplitHandler) { c = sc }); err != nil {
		fmt.Println("Error invoking container:", err)
		return nil, errors.NewDependencyNotFoundError(fmt.Sprintf(Message, "de topsecretHandler"), err)
	}
	return c, nil
}
