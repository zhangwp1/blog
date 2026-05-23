package dto

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type UpdateTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
