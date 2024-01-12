package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(Id string) (*entity.Category, error) {
	category := service.Repository.FindById(Id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}