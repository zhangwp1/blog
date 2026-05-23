package dto

import "time"

type CreateCommentRequest struct {
	ArticleID     uint   `json:"article_id" binding:"required"`
	ParentID      uint   `json:"parent_id"`
	AuthorName    string `json:"author_name" binding:"required"`
	AuthorEmail   string `json:"author_email"`
	AuthorWebsite string `json:"author_website"`
	Content       string `json:"content" binding:"required"`
}

type AdminReplyRequest struct {
	Content string `json:"content" binding:"required"`
}

type CommentResponse struct {
	ID            uint              `json:"id"`
	ArticleID     uint              `json:"article_id"`
	ParentID      uint              `json:"parent_id"`
	AuthorName    string            `json:"author_name"`
	AuthorEmail   string            `json:"author_email,omitempty"`
	AuthorWebsite string            `json:"author_website,omitempty"`
	Content       string            `json:"content"`
	IsApproved    int8              `json:"is_approved"`
	IsAdmin       bool              `json:"is_admin"`
	CreatedAt     time.Time         `json:"created_at"`
	Children      []*CommentResponse `json:"children,omitempty"`
}

type CommentListQuery struct {
	Page       int  `form:"page"`
	PageSize   int  `form:"page_size"`
	IsApproved *int `form:"is_approved"`
}
