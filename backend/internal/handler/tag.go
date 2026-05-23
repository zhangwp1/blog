package handler

import (
	"blog/internal/dto"
	"blog/internal/service"
	"blog/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler(tagService *service.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) List(c *gin.Context) {
	tags, err := h.tagService.FindAll()
	if err != nil {
		utils.Error(c, 500, "查询失败")
		return
	}
	utils.Success(c, tags)
}

func (h *TagHandler) Create(c *gin.Context) {
	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	tag, err := h.tagService.Create(req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, tag)
}

func (h *TagHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	tag, err := h.tagService.Update(uint(id), req)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, tag)
}

func (h *TagHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}

	if err := h.tagService.Delete(uint(id)); err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, nil)
}
