package usecase

import (
	entity "starshipCommsResolver/internal/pkg/entity"
	entity2 "starshipCommsResolver/internal/pkg/response"

	"github.com/stretchr/testify/mock"
)

type TopsecrestUseCaseMock struct {
	mock.Mock
}

func NewTopsecrestUseCaseMock() *TopsecrestUseCaseMock {
	return &TopsecrestUseCaseMock{}
}

func (m *TopsecrestUseCaseMock) CreateTopSecret(satellite *entity.SatelliteRequest) (*entity2.Response, error) {
	args := m.Called(satellite)
	return args.Get(0).(*entity2.Response), args.Error(1)
}
