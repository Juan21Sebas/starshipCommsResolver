package di

import (
	"testing"

	mock "starshipCommsResolver/internal/infrastructure/di/mock"
	gserrors "starshipCommsResolver/internal/infrastructure/errors"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec_BuildContainerAll(t *testing.T) {
	Convey("BuildContainer", t, func() {
		mock := &mock.DependencyInjectionMock{}
		Convey("provideServer", func() {
			Convey("Error al inyectar server en el contenedor", func() {
				mock.Err = true
				mock.Name = "Server"
				err := provideServer(mock)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &gserrors.InvalidDependencyInjectionProvideError{})
			})
			Convey("Server injectado correctamente en el contenedor", func() {
				mock.Err = false
				err := provideServer(mock)
				So(err, ShouldBeNil)
			})
		})
	})
}
