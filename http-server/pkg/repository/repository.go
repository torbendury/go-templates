package repository

import (
	"errors"

	"github.com/torbendury/go-templates/http-server/pkg/model"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type UserRepository interface {
	Get(*model.User) (*model.User, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
}
