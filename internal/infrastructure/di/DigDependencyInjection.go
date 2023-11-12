package di

import (
	"fmt"

	controllers "starshipCommsResolver/internal/infrastructure/api/starshipCommsResolver"
	server "starshipCommsResolver/internal/infrastructure/api/starshipCommsResolver"
	gserrors "starshipCommsResolver/internal/infrastructure/errors"
	"starshipCommsResolver/internal/pkg/ports"
	"starshipCommsResolver/internal/pkg/ports/di"
	services "starshipCommsResolver/internal/pkg/service"
	caseUse "starshipCommsResolver/internal/pkg/usecase"
	//di "go.uber.org/dig"
)

const Message = "No es posible proveer una instancia de %s"

type DigDependencyInjection struct {
	Container di.DependencyInjection
}

// compila las dependencias de cada hanldler a utlizar en este caso
func (di *DigDependencyInjection) BuildContainer() (di.DependencyInjection, error) {

	if err := provideServer(di.Container); err != nil {
		return di.Container, err
	}

	if err := provideTopSecret(di.Container); err != nil {
		return di.Container, err
	}
	if err := provideTopSecretSplit(di.Container); err != nil {
		return di.Container, err
	}
	if err := providerport(di.Container); err != nil {
		return di.Container, err
	}
	return di.Container, nil
}

// Se configura un provider para topsecret y las dependencias necesarias para este hanlder
func provideTopSecret(container di.DependencyInjection) error {
	if err := container.Provide(services.NewService); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "NewService"), err)
	}
	if err := container.Provide(caseUse.NewTopsecretUseCase); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "NewTopsecretUseCase"), err)
	}
	if err := container.Provide(controllers.NewHandler); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "topsecretHandler"), err)
	}
	return nil
}

func providerport(container di.DependencyInjection) error {
	container.Provide(func() ports.CommunicationServices { return &services.Service{} })
	return nil
}

func provideServer(container di.DependencyInjection) error {
	if err := container.Provide(server.NewServer); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "Server"), err)
	}
	return nil
}

// Se configura un provider para TopSecretSplit y las dependencias necesarias para este hanlder
func provideTopSecretSplit(container di.DependencyInjection) error {
	if err := container.Provide(controllers.NewSplitHandler); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "topsecretHandler"), err)
	}
	if err := container.Provide(caseUse.NewTopsecretSplitUseCase); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "NewTopsecretUseCase"), err)
	}

	if err := container.Provide(services.NewServiceSplit); err != nil {
		return gserrors.NewInvalidDependencyInjectionProvideError(
			fmt.Sprintf(Message, "NewService"), err)
	}

	return nil
}
