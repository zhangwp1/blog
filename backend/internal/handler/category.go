package handler

import (
	"blog/internal/dto"
	"blog/internal/service"
	"blog/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.categoryService.FindAll()
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}
	utils.Success(c, categories)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	category, err := h.categoryService.Create(req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	category, err := h.categoryService.Update(uint(id), req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.categoryService.Delete(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}
