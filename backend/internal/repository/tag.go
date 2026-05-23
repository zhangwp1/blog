package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) FindAll() ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Order("id ASC").Find(&tags).Error
	return tags, err
}

func (r *TagRepository) FindByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) FindBySlug(slug string) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.Where("slug = ?", slug).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) FindByIDs(ids []uint) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

func (r *TagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

func (r *TagRepository) Count() int64 {
	var count int64
	r.db.Model(&model.Tag{}).Count(&count)
	return count
}
