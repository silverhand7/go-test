package repository

import "go-test/entity"

type CategoryRepository interface {
	FindById(id int) *entity.Category
}
