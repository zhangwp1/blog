package handler

import (
	"blog/internal/dto"
	"blog/internal/service"
	"blog/internal/utils"
	"blog/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
	articleRepo    *repository.ArticleRepository
}

func NewCommentHandler(commentService *service.CommentService, articleRepo *repository.ArticleRepository) *CommentHandler {
	return &CommentHandler{commentService: commentService, articleRepo: articleRepo}
}

func (h *CommentHandler) ListByArticle(c *gin.Context) {
	slug := c.Param("slug")
	article, err := h.articleRepo.FindBySlug(slug)
	if err != nil {
		utils.Error(c, 404, "文章不存在")
		return
	}

	comments, err := h.commentService.GetByArticleSlug(article.ID)
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}
	utils.Success(c, comments)
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	comment, err := h.commentService.Create(req, ip, userAgent)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, comment)
}

func (h *CommentHandler) ListAdmin(c *gin.Context) {
	var query dto.CommentListQuery
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

	comments, total, err := h.commentService.List(query)
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}
	utils.PageSuccess(c, comments, total, query.Page, query.PageSize)
}

func (h *CommentHandler) Approve(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.commentService.Approve(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}

func (h *CommentHandler) Reject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.commentService.Reject(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}

func (h *CommentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.commentService.Delete(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}

func (h *CommentHandler) AdminReply(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	var req dto.AdminReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	comment, err := h.commentService.FindByID(uint(id))
	if err != nil {
		utils.Error(c, 404, "评论不存在")
		return
	}

	result, err := h.commentService.AdminReply(comment.ArticleID, uint(id), req.Content)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, result)
}
