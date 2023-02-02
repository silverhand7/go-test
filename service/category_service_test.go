package service

import (
	"go-test/entity"
	"go-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositorMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_FirstNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", 1).Return(nil)

	category, err := categoryService.First(1)

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_FirstFound(t *testing.T) {
	category := entity.Category{
		Id:   1,
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", 1).Return(category)

	result, err := categoryService.First(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
