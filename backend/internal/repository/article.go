package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticleRepository) Update(article *model.Article) error {
	if err := r.db.Save(article).Error; err != nil {
		return err
	}
	return r.db.Model(article).Association("Tags").Replace(article.Tags)
}

func (r *ArticleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

func (r *ArticleRepository) FindByID(id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Category").Preload("Tags").Preload("Author").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) FindBySlug(slug string) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Category").Preload("Tags").Preload("Author").
		Where("slug = ?", slug).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

type ArticleFilter struct {
	IsPublished *bool
	CategoryID  uint
	TagSlug     string
	Keyword     string
	Year        int
	Month       int
	Pinned      *bool
	Page        int
	PageSize    int
}

func (r *ArticleRepository) FindList(filter ArticleFilter) ([]model.Article, int64, error) {
	query := r.db.Model(&model.Article{}).Preload("Category").Preload("Tags")

	if filter.IsPublished != nil {
		query = query.Where("is_published = ?", *filter.IsPublished)
	}
	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.TagSlug != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Joins("JOIN tags ON tags.id = article_tags.tag_id").
			Where("tags.slug = ?", filter.TagSlug)
	}
	if filter.Keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}
	if filter.Year > 0 {
		if filter.Month > 0 {
			query = query.Where("YEAR(published_at) = ? AND MONTH(published_at) = ?", filter.Year, filter.Month)
		} else {
			query = query.Where("YEAR(published_at) = ?", filter.Year)
		}
	}
	if filter.Pinned != nil {
		query = query.Where("pinned = ?", *filter.Pinned)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.PageSize <= 0 {
		filter.PageSize = 10
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}

	var articles []model.Article
	err := query.Order("pinned DESC, published_at DESC").
		Offset((filter.Page - 1) * filter.PageSize).
		Limit(filter.PageSize).
		Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (r *ArticleRepository) FindPrevNext(slug string) (*model.Article, *model.Article, error) {
	current, err := r.FindBySlug(slug)
	if err != nil {
		return nil, nil, err
	}

	var prev, next model.Article
	r.db.Where("is_published = ? AND published_at > ?", true, current.PublishedAt).
		Order("published_at ASC").First(&prev)
	r.db.Where("is_published = ? AND published_at < ?", true, current.PublishedAt).
		Order("published_at DESC").First(&next)

	var prevPtr, nextPtr *model.Article
	if prev.ID != 0 {
		prevPtr = &prev
	}
	if next.ID != 0 {
		nextPtr = &next
	}
	return prevPtr, nextPtr, nil
}

func (r *ArticleRepository) IncrementViewCount(id uint) {
	r.db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1"))
}

func (r *ArticleRepository) Count() int64 {
	var count int64
	r.db.Model(&model.Article{}).Count(&count)
	return count
}

func (r *ArticleRepository) CountByStatus(isPublished bool) int64 {
	var count int64
	r.db.Model(&model.Article{}).Where("is_published = ?", isPublished).Count(&count)
	return count
}
