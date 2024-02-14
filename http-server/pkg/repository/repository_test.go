package repository

import "github.com/torbendury/go-templates/http-server/pkg/model"

type UserRepositoryMock struct {
	Users []model.User
}

func (urm *UserRepositoryMock) Get(*model.User) (*model.User, error) {
	return nil, nil
}

func (urm *UserRepositoryMock) Create(*model.User) (*model.User, error) {
	return nil, nil
}

func (urm *UserRepositoryMock) Update(*model.User) (*model.User, error) {
	return nil, nil
}

func (urm *UserRepositoryMock) Delete(*model.User) error {
	return nil
}
