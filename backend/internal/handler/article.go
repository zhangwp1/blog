package handler

import (
	"blog/internal/dto"
	"blog/internal/service"
	"blog/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *service.ArticleService
}

func NewArticleHandler(articleService *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) ListPublic(c *gin.Context) {
	var query dto.ArticleListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	articles, total, err := h.articleService.ListPublic(query)
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}

	utils.PageSuccess(c, articles, total, query.Page, query.PageSize)
}

func (h *ArticleHandler) ListAdmin(c *gin.Context) {
	var query dto.ArticleListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	articles, total, err := h.articleService.ListAdmin(query)
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}

	utils.PageSuccess(c, articles, total, query.Page, query.PageSize)
}

func (h *ArticleHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		utils.Error(c, 404, err.Error())
		return
	}
	utils.Success(c, article)
}

func (h *ArticleHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}
	article, err := h.articleService.GetByID(uint(id))
	if err != nil {
		utils.Error(c, 404, err.Error())
		return
	}
	utils.Success(c, article)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	article, err := h.articleService.Create(userID, req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, article)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	var req dto.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	article, err := h.articleService.Update(uint(id), req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, article)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.articleService.Delete(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}

func (h *ArticleHandler) Dashboard(c *gin.Context) {
	data := h.articleService.GetDashboard()
	utils.Success(c, data)
}
