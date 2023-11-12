package di

import "go.uber.org/dig"

type DependencyInjection interface {
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
}
