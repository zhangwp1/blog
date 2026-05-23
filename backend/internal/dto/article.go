package dto

import "time"

type CreateArticleRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary"`
	CoverImage  string `json:"cover_image"`
	IsPublished *bool  `json:"is_published"`
	Pinned      *bool  `json:"pinned"`
	CategoryID  uint   `json:"category_id"`
	TagIDs      []uint `json:"tag_ids"`
}

type UpdateArticleRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Summary     string `json:"summary"`
	CoverImage  string `json:"cover_image"`
	IsPublished *bool  `json:"is_published"`
	Pinned      *bool  `json:"pinned"`
	CategoryID  *uint  `json:"category_id"`
	TagIDs      []uint `json:"tag_ids"`
}

type ArticleListQuery struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	Category  string `form:"category"`
	Tag       string `form:"tag"`
	Keyword   string `form:"keyword"`
	Year      int    `form:"year"`
	Month     int    `form:"month"`
}

type ArticleResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content,omitempty"`
	Summary     string     `json:"summary"`
	CoverImage  string     `json:"cover_image"`
	IsPublished bool       `json:"is_published"`
	Pinned      bool       `json:"pinned"`
	ViewCount   int        `json:"view_count"`
	CategoryID  uint       `json:"category_id"`
	AuthorID    uint       `json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Category    *CategoryInfo `json:"category,omitempty"`
	Author      *AuthorInfo   `json:"author,omitempty"`
	Tags        []TagInfo     `json:"tags,omitempty"`
}

type CategoryInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type AuthorInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type TagInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ArticleListResponse struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	Summary     string        `json:"summary"`
	CoverImage  string        `json:"cover_image"`
	IsPublished bool          `json:"is_published"`
	Pinned      bool          `json:"pinned"`
	ViewCount   int           `json:"view_count"`
	CategoryID  uint          `json:"category_id"`
	PublishedAt *time.Time    `json:"published_at"`
	CreatedAt   time.Time     `json:"created_at"`
	Category    *CategoryInfo `json:"category,omitempty"`
	Tags        []TagInfo     `json:"tags,omitempty"`
}
