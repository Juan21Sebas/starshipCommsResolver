package usecase

import (
	"fmt"

	entity "starshipCommsResolver/internal/pkg/entity"
	"starshipCommsResolver/internal/pkg/ports"
	entity2 "starshipCommsResolver/internal/pkg/response"
)

type TopsecrestUseCase struct {
	topsecrestService ports.CommunicationServices
}

func NewTopsecretUseCase(service ports.CommunicationServices) *TopsecrestUseCase {
	return &TopsecrestUseCase{
		topsecrestService: service,
	}
}

func (uc *TopsecrestUseCase) CreateTopSecret(satellite *entity.SatelliteRequest) (*entity2.Response, error) {

	var Position entity2.Position
	var message [][]string
	distance := []float32{}

	for _, satellite := range satellite.Satellites {
		message = append(message, satellite.Message)
		distance = append(distance, float32(satellite.Distance))
	}
	Position.Message = uc.topsecrestService.GetMessage(message...)

	Position.X, Position.Y = uc.topsecrestService.GetLocation(distance...)
	if Position.X == 0 && Position.Y == 0 || Position.Message == "" {
		//return nil, fmt.Errorf("the position or the message cannot be determined", "error")
		return nil, fmt.Errorf("the position or the message cannot be determined")

	} else {
		return &entity2.Response{
			Position: entity2.Position{
				X:       Position.X,
				Y:       Position.Y,
				Message: Position.Message,
			},
			Result: entity2.Result{
				Source:  "Find position to message",
				Details: nil,
			},
		}, nil
	}
}
