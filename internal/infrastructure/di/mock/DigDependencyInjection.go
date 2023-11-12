package mock

import (
	"errors"
	"reflect"

	"go.uber.org/dig"
)

type DependencyInjectionMock struct {
	Err  bool
	Name string
}

func (dim *DependencyInjectionMock) Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	t := reflect.TypeOf(constructor).Out(0)
	h := t.Elem().Name()
	if dim.Err && h == dim.Name {
		return errors.New("Error")
	}
	return nil
}
