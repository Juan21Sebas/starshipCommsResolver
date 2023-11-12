package usecase

import (
	"errors"
	"fmt"

	entity "starshipCommsResolver/internal/pkg/entity"
	"starshipCommsResolver/internal/pkg/ports"
	entity2 "starshipCommsResolver/internal/pkg/response"
)

var satelliteDataMap = make(map[string]*entity.SatelliteData)

type TopsecretSplitUseCase struct {
	topsecretsplitService ports.CommunicationServices
}

func NewTopsecretSplitUseCase(service ports.CommunicationServices) *TopsecretSplitUseCase {
	return &TopsecretSplitUseCase{}
}

func (uc *TopsecretSplitUseCase) CreateTopSecretSplit(satelliteName *string, satelliteData *entity.SatelliteData) (*entity2.Response, error) {

	satelliteDataMap[*satelliteName] = satelliteData
	count := len(satelliteDataMap)
	if count > 3 {
		return nil, errors.New("There are enough positions to track the message and the location of the satellites.")
	}

	return &entity2.Response{
		Result: entity2.Result{
			Source:  "Message and distance added successfully",
			Details: nil,
		},
	}, nil

}

func (uc *TopsecretSplitUseCase) GetAll(satelliteName string, satellite *entity.Satellitequery) (*entity2.Response, error) {
	//var resp entity2.Response
	var Position entity2.Position
	_, exists := satelliteDataMap[satelliteName]
	if !exists {
		return nil, errors.New("Satellite data not found")
	}
	if len(satelliteDataMap) < 3 {
		return nil, errors.New("Insufficient satellite positions")
	}
	var allMessages [][]string
	for _, data := range satelliteDataMap {
		allMessages = append(allMessages, data.Message)
	}

	Position.Message = uc.topsecretsplitService.GetMessage(allMessages...)

	alllocation := []float32{}
	for _, data := range satelliteDataMap {
		alllocation = append(alllocation, data.Distance)
	}
	Position.X, Position.Y = uc.topsecretsplitService.GetLocation(alllocation...)
	if Position.X == 0 && Position.Y == 0 || Position.Message == "" {
		return nil, errors.New("Insufficient satellite positions")
	} else {
		satelliteDataMap = nil
		satelliteDataMap = make(map[string]*entity.SatelliteData)
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

func (uc *TopsecretSplitUseCase) PatchtopsecretSplit(satelliteName string, satelliteData *entity.SatelliteData) (*entity2.Response, error) {
	//var resp entity2.Response
	//var Position entity2.Position
	existingData, exists := satelliteDataMap[satelliteName]
	if exists {
		// Reemplazar los datos existentes con los nuevos valores
		existingData = satelliteData
		satelliteDataMap[satelliteName] = existingData
		//fmt.Println("Data updated.:", satelliteDataMap[satelliteName])
	} else {
		fmt.Println("Satellite not found in the map..")
	}
	return &entity2.Response{
		Result: entity2.Result{
			Source:  "Data updated.",
			Details: nil,
		},
	}, nil
}
