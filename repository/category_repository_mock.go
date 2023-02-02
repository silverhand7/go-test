package repository

import (
	"go-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositorMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositorMock) FindById(id int) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
