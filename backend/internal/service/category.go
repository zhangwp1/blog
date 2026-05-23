package service

import (
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) FindAll() ([]model.Category, error) {
	return s.categoryRepo.FindAll()
}

func (s *CategoryService) Create(req dto.CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{
		Name:      req.Name,
		Slug:      req.Slug,
		SortOrder: req.SortOrder,
	}
	if err := s.categoryRepo.Create(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *CategoryService) Update(id uint, req dto.UpdateCategoryRequest) (*model.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("分类不存在")
		}
		return nil, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Slug != "" {
		category.Slug = req.Slug
	}
	if req.SortOrder != nil {
		category.SortOrder = *req.SortOrder
	}

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Delete(id uint) error {
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("分类不存在")
		}
		return err
	}
	return s.categoryRepo.Delete(id)
}
