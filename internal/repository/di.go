package repository

import "go.uber.org/dig"

type (
	Holder struct {
		dig.In
		Post Post
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewPost); err != nil {
		return err
	}

	return nil
}
