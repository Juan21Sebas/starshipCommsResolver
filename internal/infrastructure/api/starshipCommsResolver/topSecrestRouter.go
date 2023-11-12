package starshipcommsresolver

import (
	"fmt"
	"starshipCommsResolver/internal/infrastructure/errors"

	"github.com/gin-gonic/gin"

	"go.uber.org/dig"
)

const Message = "No es posible obtener una instancia %s"

func (server *Server) TopSecretRoutes(router *gin.Engine) error {
	c, err := getTopSecretHandler(server.container)
	if err != nil {
		return err
	}
	routes := router
	{
		routes.POST("/topsecret/", c.topsecret)
		//e.POST("/topsecret/", topsecretHandler.topsecret)
	}
	return nil
}

func getTopSecretHandler(container *dig.Container) (*topsecretHandler, error) {
	var c *topsecretHandler

	if err := container.Invoke(func(sc *topsecretHandler) { c = sc }); err != nil {
		fmt.Println("Error invoking container:", err)
		return nil, errors.NewDependencyNotFoundError(fmt.Sprintf(Message, "de topsecretHandler"), err)
	}
	return c, nil
}
