package usecase_test

import (
	"testing"

	mock1 "starshipCommsResolver/internal/infrastructure/api/starshipCommsResolver/mock"
	"starshipCommsResolver/internal/pkg/entity"
	"starshipCommsResolver/internal/pkg/response"
	"starshipCommsResolver/internal/pkg/usecase"
	mock2 "starshipCommsResolver/internal/pkg/usecase/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTopsecrestUseCase_CreateTopSecret_Success(t *testing.T) {
	mockService := new(mock1.CommunicationServicesMock)
	usecase := usecase.NewTopsecretUseCase(mockService)
	caseUsemock := mock2.NewTopsecrestUseCaseMock()

	satelliteRequest := &entity.SatelliteRequest{
		Satellites: []entity.Satellite{
			{
				Name:     "kenobi",
				Distance: 100,
				Message:  []string{"mensaje1"},
			},
		},
	}

	mockService.On("GetLocation", mock.Anything).Return(float32(300.0), float32(400.0))
	mockService.On("GetMessage", mock.Anything).Return("Mensaje decodificado")

	expectedResponse := &response.Response{
		Position: response.Position{
			X:       300.0,
			Y:       400.0,
			Message: "Mensaje decodificado",
		},
		Result: response.Result{
			Source:  "Find position to message",
			Details: nil,
		},
	}

	caseUsemock.On("CreateTopSecret", satelliteRequest).Return(expectedResponse, nil)

	response, _ := caseUsemock.CreateTopSecret(satelliteRequest)
	expectresponseuse, _ := usecase.CreateTopSecret(satelliteRequest)

	assert.Equal(t, expectresponseuse, response)
}
