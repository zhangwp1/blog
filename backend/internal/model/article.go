package model

import "time"

type Article struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Slug        string    `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Content     string    `gorm:"type:mediumtext;not null" json:"content"`
	Summary     string    `gorm:"size:512" json:"summary"`
	CoverImage  string    `gorm:"size:255" json:"cover_image"`
	IsPublished bool      `gorm:"default:false" json:"is_published"`
	Pinned      bool      `gorm:"default:false" json:"pinned"`
	ViewCount   int       `gorm:"default:0" json:"view_count"`
	CategoryID  uint      `gorm:"default:0" json:"category_id"`
	AuthorID    uint      `gorm:"not null" json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Author   User     `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Tags     []Tag    `gorm:"many2many:article_tags" json:"tags,omitempty"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}
