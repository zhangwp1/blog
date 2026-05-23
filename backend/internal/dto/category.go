package dto

type CreateCategoryRequest struct {
	Name      string `json:"name" binding:"required"`
	Slug      string `json:"slug" binding:"required"`
	SortOrder int    `json:"sort_order"`
}

type UpdateCategoryRequest struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	SortOrder *int   `json:"sort_order"`
}
