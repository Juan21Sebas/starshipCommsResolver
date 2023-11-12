package main

import (
	"os"
	server "starshipCommsResolver/internal/infrastructure/api/starshipCommsResolver"
	"starshipCommsResolver/internal/infrastructure/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {

	if err := bootstrapping(); err != nil {
		os.Exit(-1)
	}
}

func bootstrapping() error {
	g := gin.Default()
	digDependencyInjection := &di.DigDependencyInjection{Container: dig.New()}
	d, err := digDependencyInjection.BuildContainer()
	if err != nil {
		return err
	}
	//fmt.Println(d.String())

	srv := server.NewServer(g, d.(*dig.Container))
	if err := srv.MapRoutes(); err != nil {
		return err
	}
	return srv.Start()
}
