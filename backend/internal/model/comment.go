package model

import "time"

type Comment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ArticleID     uint      `gorm:"not null;index" json:"article_id"`
	ParentID      uint      `gorm:"default:0" json:"parent_id"`
	AuthorName    string    `gorm:"size:64;not null" json:"author_name"`
	AuthorEmail   string    `gorm:"size:128" json:"author_email"`
	AuthorWebsite string    `gorm:"size:255" json:"author_website"`
	Content       string    `gorm:"type:text;not null" json:"content"`
	IsApproved    int8      `gorm:"default:0" json:"is_approved"`
	IsAdmin       bool      `gorm:"default:false" json:"is_admin"`
	IP            string    `gorm:"size:45" json:"-"`
	UserAgent     string    `gorm:"size:255" json:"-"`
	CreatedAt     time.Time `json:"created_at"`

	Children []*Comment `gorm:"-" json:"children,omitempty"`
}
