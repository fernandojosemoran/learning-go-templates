package user

import (
	"errors"
	"strings"
)

type userDto struct {
	Name string
}

func CreateUserDto(name string) (*userDto, error) {
	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return nil, errors.New("Name property is required.")
	}

	return &userDto{
			Name: name,
		},
		nil
}
