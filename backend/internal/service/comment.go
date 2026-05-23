package service

import (
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

func NewCommentService(commentRepo *repository.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) GetByArticleSlug(articleID uint) ([]dto.CommentResponse, error) {
	comments, err := s.commentRepo.FindByArticleID(articleID)
	if err != nil {
		return nil, err
	}
	return buildCommentTree(comments), nil
}

func (s *CommentService) List(query dto.CommentListQuery) ([]dto.CommentResponse, int64, error) {
	filter := repository.CommentListFilter{
		IsApproved: (*int)(query.IsApproved),
		Page:       query.Page,
		PageSize:   query.PageSize,
	}

	comments, total, err := s.commentRepo.FindList(filter)
	if err != nil {
		return nil, 0, err
	}

	result := make([]dto.CommentResponse, len(comments))
	for i, c := range comments {
		result[i] = toCommentResponse(&c)
	}
	return result, total, nil
}

func (s *CommentService) Create(req dto.CreateCommentRequest, ip, userAgent string) (*dto.CommentResponse, error) {
	comment := model.Comment{
		ArticleID:     req.ArticleID,
		ParentID:      req.ParentID,
		AuthorName:    req.AuthorName,
		AuthorEmail:   req.AuthorEmail,
		AuthorWebsite: req.AuthorWebsite,
		Content:       req.Content,
		IP:            ip,
		UserAgent:     userAgent,
	}

	if err := s.commentRepo.Create(&comment); err != nil {
		return nil, err
	}

	resp := toCommentResponse(&comment)
	return &resp, nil
}

func (s *CommentService) FindByID(id uint) (*dto.CommentResponse, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("评论不存在")
		}
		return nil, err
	}
	resp := toCommentResponse(comment)
	return &resp, nil
}

func (s *CommentService) Approve(id uint) error {
	_, err := s.commentRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	return s.commentRepo.UpdateStatus(id, 1)
}

func (s *CommentService) Reject(id uint) error {
	_, err := s.commentRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	return s.commentRepo.UpdateStatus(id, 2)
}

func (s *CommentService) Delete(id uint) error {
	_, err := s.commentRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	return s.commentRepo.Delete(id)
}

func (s *CommentService) AdminReply(articleID, parentID uint, content string) (*dto.CommentResponse, error) {
	comment := model.Comment{
		ArticleID:  articleID,
		ParentID:   parentID,
		AuthorName: "博主",
		Content:    content,
		IsApproved: 1,
		IsAdmin:    true,
	}

	if err := s.commentRepo.Create(&comment); err != nil {
		return nil, err
	}

	resp := toCommentResponse(&comment)
	return &resp, nil
}

func (s *CommentService) GetDashboard() map[string]interface{} {
	return map[string]interface{}{
		"total_comments":   s.commentRepo.Count(),
		"pending_comments": s.commentRepo.CountByStatus(0),
		"approved_comments": s.commentRepo.CountByStatus(1),
	}
}

func buildCommentTree(comments []model.Comment) []dto.CommentResponse {
	nodeMap := make(map[uint]*dto.CommentResponse)
	var roots []dto.CommentResponse

	for _, c := range comments {
		resp := toCommentResponse(&c)
		nodeMap[c.ID] = &resp
	}

	for _, c := range comments {
		node := nodeMap[c.ID]
		if c.ParentID == 0 {
			roots = append(roots, *node)
		} else {
			parent, ok := nodeMap[c.ParentID]
			if ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return roots
}

func toCommentResponse(c *model.Comment) dto.CommentResponse {
	return dto.CommentResponse{
		ID:            c.ID,
		ArticleID:     c.ArticleID,
		ParentID:      c.ParentID,
		AuthorName:    c.AuthorName,
		AuthorEmail:   c.AuthorEmail,
		AuthorWebsite: c.AuthorWebsite,
		Content:       c.Content,
		IsApproved:    c.IsApproved,
		IsAdmin:       c.IsAdmin,
		CreatedAt:     c.CreatedAt,
	}
}
