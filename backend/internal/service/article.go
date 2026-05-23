package service

import (
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/repository"
	"errors"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ArticleService struct {
	articleRepo  *repository.ArticleRepository
	categoryRepo *repository.CategoryRepository
	tagRepo      *repository.TagRepository
}

func NewArticleService(
	articleRepo *repository.ArticleRepository,
	categoryRepo *repository.CategoryRepository,
	tagRepo *repository.TagRepository,
) *ArticleService {
	return &ArticleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (s *ArticleService) ListPublic(query dto.ArticleListQuery) ([]dto.ArticleListResponse, int64, error) {
	filter := repository.ArticleFilter{
		IsPublished: boolPtr(true),
		Keyword:     query.Keyword,
		Year:        query.Year,
		Month:       query.Month,
		Page:        query.Page,
		PageSize:    query.PageSize,
	}

	if query.Category != "" {
		cat, err := s.categoryRepo.FindBySlug(query.Category)
		if err == nil {
			filter.CategoryID = cat.ID
		}
	}
	if query.Tag != "" {
		filter.TagSlug = query.Tag
	}

	articles, total, err := s.articleRepo.FindList(filter)
	if err != nil {
		return nil, 0, err
	}

	result := make([]dto.ArticleListResponse, len(articles))
	for i, a := range articles {
		result[i] = toArticleListResponse(&a)
	}
	return result, total, nil
}

func (s *ArticleService) ListAdmin(query dto.ArticleListQuery) ([]dto.ArticleListResponse, int64, error) {
	filter := repository.ArticleFilter{
		Keyword:  query.Keyword,
		Page:     query.Page,
		PageSize: query.PageSize,
	}

	articles, total, err := s.articleRepo.FindList(filter)
	if err != nil {
		return nil, 0, err
	}

	result := make([]dto.ArticleListResponse, len(articles))
	for i, a := range articles {
		result[i] = toArticleListResponse(&a)
	}
	return result, total, nil
}

func (s *ArticleService) GetBySlug(slug string) (*dto.ArticleResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}
	s.articleRepo.IncrementViewCount(article.ID)
	resp := toArticleResponse(article)
	return &resp, nil
}

func (s *ArticleService) GetByID(id uint) (*dto.ArticleResponse, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}
	resp := toArticleResponse(article)
	return &resp, nil
}

func (s *ArticleService) Create(userID uint, req dto.CreateArticleRequest) (*dto.ArticleResponse, error) {
	article := model.Article{
		Title:      req.Title,
		Slug:       slug.Make(req.Title),
		Content:    req.Content,
		Summary:    req.Summary,
		CoverImage: req.CoverImage,
		AuthorID:   userID,
		CategoryID: req.CategoryID,
	}
	if req.IsPublished != nil && *req.IsPublished {
		article.IsPublished = true
		now := time.Now()
		article.PublishedAt = &now
	}
	if req.Pinned != nil {
		article.Pinned = *req.Pinned
	}
	if len(req.TagIDs) > 0 {
		tags, err := s.tagRepo.FindByIDs(req.TagIDs)
		if err != nil {
			return nil, err
		}
		article.Tags = tags
	}

	if err := s.articleRepo.Create(&article); err != nil {
		return nil, err
	}

	return s.GetByID(article.ID)
}

func (s *ArticleService) Update(id uint, req dto.UpdateArticleRequest) (*dto.ArticleResponse, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}

	if req.Title != "" {
		article.Title = req.Title
		article.Slug = slug.Make(req.Title)
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	article.CoverImage = req.CoverImage
	if req.IsPublished != nil {
		if *req.IsPublished && !article.IsPublished {
			now := time.Now()
			article.PublishedAt = &now
		}
		article.IsPublished = *req.IsPublished
	}
	if req.Pinned != nil {
		article.Pinned = *req.Pinned
	}
	if req.CategoryID != nil {
		article.CategoryID = *req.CategoryID
	}
	if req.TagIDs != nil {
		tags, err := s.tagRepo.FindByIDs(req.TagIDs)
		if err != nil {
			return nil, err
		}
		article.Tags = tags
	}

	if err := s.articleRepo.Update(article); err != nil {
		return nil, err
	}

	return s.GetByID(article.ID)
}

func (s *ArticleService) Delete(id uint) error {
	_, err := s.articleRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		return err
	}
	return s.articleRepo.Delete(id)
}

func (s *ArticleService) GetDashboard() map[string]interface{} {
	return map[string]interface{}{
		"total_articles":       s.articleRepo.Count(),
		"published_articles":   s.articleRepo.CountByStatus(true),
		"draft_articles":       s.articleRepo.CountByStatus(false),
	}
}

func toArticleResponse(a *model.Article) dto.ArticleResponse {
	resp := dto.ArticleResponse{
		ID:          a.ID,
		Title:       a.Title,
		Slug:        a.Slug,
		Content:     a.Content,
		Summary:     a.Summary,
		CoverImage:  a.CoverImage,
		IsPublished: a.IsPublished,
		Pinned:      a.Pinned,
		ViewCount:   a.ViewCount,
		CategoryID:  a.CategoryID,
		AuthorID:    a.AuthorID,
		PublishedAt: a.PublishedAt,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}

	if a.Category.ID != 0 {
		resp.Category = &dto.CategoryInfo{
			ID:   a.Category.ID,
			Name: a.Category.Name,
			Slug: a.Category.Slug,
		}
	}

	if a.Author.ID != 0 {
		resp.Author = &dto.AuthorInfo{
			ID:       a.Author.ID,
			Username: a.Author.Username,
			Nickname: a.Author.Nickname,
			Avatar:   a.Author.Avatar,
		}
	}

	tags := make([]dto.TagInfo, len(a.Tags))
	for i, t := range a.Tags {
		tags[i] = dto.TagInfo{ID: t.ID, Name: t.Name, Slug: t.Slug}
	}
	resp.Tags = tags

	return resp
}

func toArticleListResponse(a *model.Article) dto.ArticleListResponse {
	resp := dto.ArticleListResponse{
		ID:          a.ID,
		Title:       a.Title,
		Slug:        a.Slug,
		Summary:     a.Summary,
		CoverImage:  a.CoverImage,
		IsPublished: a.IsPublished,
		Pinned:      a.Pinned,
		ViewCount:   a.ViewCount,
		CategoryID:  a.CategoryID,
		PublishedAt: a.PublishedAt,
		CreatedAt:   a.CreatedAt,
	}

	if a.Category.ID != 0 {
		resp.Category = &dto.CategoryInfo{
			ID:   a.Category.ID,
			Name: a.Category.Name,
			Slug: a.Category.Slug,
		}
	}

	tags := make([]dto.TagInfo, len(a.Tags))
	for i, t := range a.Tags {
		tags[i] = dto.TagInfo{ID: t.ID, Name: t.Name, Slug: t.Slug}
	}
	resp.Tags = tags

	return resp
}

func boolPtr(b bool) *bool {
	return &b
}
