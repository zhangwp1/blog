package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) FindByArticleID(articleID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("article_id = ? AND is_approved = ?", articleID, 1).
		Order("created_at ASC").Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) FindList(query CommentListFilter) ([]model.Comment, int64, error) {
	q := r.db.Model(&model.Comment{})

	if query.IsApproved != nil {
		q = q.Where("is_approved = ?", *query.IsApproved)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.Page <= 0 {
		query.Page = 1
	}

	var comments []model.Comment
	err := q.Order("created_at DESC").Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).Find(&comments).Error
	return comments, total, err
}

type CommentListFilter struct {
	IsApproved *int
	Page       int
	PageSize   int
}

func (r *CommentRepository) FindByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) UpdateStatus(id uint, status int8) error {
	return r.db.Model(&model.Comment{}).Where("id = ?", id).Update("is_approved", status).Error
}

func (r *CommentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Comment{}, id).Error
}

func (r *CommentRepository) CountByStatus(isApproved int8) int64 {
	var count int64
	r.db.Model(&model.Comment{}).Where("is_approved = ?", isApproved).Count(&count)
	return count
}

func (r *CommentRepository) Count() int64 {
	var count int64
	r.db.Model(&model.Comment{}).Count(&count)
	return count
}
