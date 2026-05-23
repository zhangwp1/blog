package handler

import (
	"blog/internal/repository"
	"blog/internal/service"

	"gorm.io/gorm"
)

type Handlers struct {
	Auth    *AuthHandler
	Article *ArticleHandler
	Category *CategoryHandler
	Tag     *TagHandler
	Comment *CommentHandler
}

func InitHandlers(db *gorm.DB) *Handlers {
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	tagRepo := repository.NewTagRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	authService := service.NewAuthService(userRepo)
	articleService := service.NewArticleService(articleRepo, categoryRepo, tagRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	tagService := service.NewTagService(tagRepo)
	commentService := service.NewCommentService(commentRepo)

	return &Handlers{
		Auth:     NewAuthHandler(authService),
		Article:  NewArticleHandler(articleService),
		Category: NewCategoryHandler(categoryService),
		Tag:      NewTagHandler(tagService),
		Comment:  NewCommentHandler(commentService, articleRepo),
	}
}
