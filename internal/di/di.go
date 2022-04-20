package di

import (
	"tix-rest-project/internal/repository"
	"tix-rest-project/internal/service"
	"tix-rest-project/internal/shared"

	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func init() {
	if err := shared.Register(Container); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewDatabase); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewCache); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewEcho); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}
}
