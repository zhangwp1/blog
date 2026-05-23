package service

import (
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type TagService struct {
	tagRepo *repository.TagRepository
}

func NewTagService(tagRepo *repository.TagRepository) *TagService {
	return &TagService{tagRepo: tagRepo}
}

func (s *TagService) FindAll() ([]model.Tag, error) {
	return s.tagRepo.FindAll()
}

func (s *TagService) Create(req dto.CreateTagRequest) (*model.Tag, error) {
	tag := model.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}
	if err := s.tagRepo.Create(&tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

func (s *TagService) Update(id uint, req dto.UpdateTagRequest) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("标签不存在")
		}
		return nil, err
	}

	if req.Name != "" {
		tag.Name = req.Name
	}
	if req.Slug != "" {
		tag.Slug = req.Slug
	}

	if err := s.tagRepo.Update(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *TagService) Delete(id uint) error {
	_, err := s.tagRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("标签不存在")
		}
		return err
	}
	return s.tagRepo.Delete(id)
}
