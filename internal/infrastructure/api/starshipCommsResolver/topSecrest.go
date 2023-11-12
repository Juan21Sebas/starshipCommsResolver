package starshipcommsresolver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"starshipCommsResolver/internal/pkg/entity"
	ports "starshipCommsResolver/internal/pkg/ports"
	entity2 "starshipCommsResolver/internal/pkg/response"
	"starshipCommsResolver/internal/pkg/usecase"
)

type topsecretHandler struct {
	topsecretUseCase *usecase.TopsecrestUseCase
}

func NewHandler(services ports.CommunicationServices) *topsecretHandler {
	fmt.Println("Creating topsecretHandler")
	orderUseCase := usecase.NewTopsecretUseCase(services)
	return &topsecretHandler{
		topsecretUseCase: orderUseCase,
	}
}

func (o *topsecretHandler) topsecret(c *gin.Context) {

	var satellite entity.SatelliteRequest
	if err := c.BindJSON(&satellite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretUseCase.CreateTopSecret(&satellite)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}
