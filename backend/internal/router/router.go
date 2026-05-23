package router

import (
	"blog/config"
	"blog/internal/handler"
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	if config.AppConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// 初始化依赖
	handlers := handler.InitHandlers(db)

	api := r.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// 公开路由组
		public := api.Group("")
		{
			public.POST("/auth/login", handlers.Auth.Login)
			public.GET("/articles", handlers.Article.ListPublic)
			public.GET("/articles/:slug", handlers.Article.GetBySlug)
			public.GET("/categories", handlers.Category.List)
			public.GET("/tags", handlers.Tag.List)
			public.GET("/articles/:slug/comments", handlers.Comment.ListByArticle)
			public.POST("/articles/:slug/comments", handlers.Comment.Create)
		}

		// 需认证路由组
		authRequired := api.Group("")
		authRequired.Use(middleware.Auth())
		{
			authRequired.GET("/auth/profile", handlers.Auth.Profile)
		}

		// 管理端路由组
		admin := api.Group("/admin")
		admin.Use(middleware.Auth())
		{
			admin.GET("/dashboard", handlers.Article.Dashboard)

			admin.GET("/articles", handlers.Article.ListAdmin)
			admin.GET("/articles/:id", handlers.Article.GetByID)
			admin.POST("/articles", handlers.Article.Create)
			admin.PUT("/articles/:id", handlers.Article.Update)
			admin.DELETE("/articles/:id", handlers.Article.Delete)

			admin.GET("/categories", handlers.Category.List)
			admin.POST("/categories", handlers.Category.Create)
			admin.PUT("/categories/:id", handlers.Category.Update)
			admin.DELETE("/categories/:id", handlers.Category.Delete)

			admin.GET("/tags", handlers.Tag.List)
			admin.POST("/tags", handlers.Tag.Create)
			admin.PUT("/tags/:id", handlers.Tag.Update)
			admin.DELETE("/tags/:id", handlers.Tag.Delete)

			admin.GET("/comments", handlers.Comment.ListAdmin)
			admin.PATCH("/comments/:id/approve", handlers.Comment.Approve)
			admin.PATCH("/comments/:id/reject", handlers.Comment.Reject)
			admin.DELETE("/comments/:id", handlers.Comment.Delete)
			admin.POST("/comments/:id/reply", handlers.Comment.AdminReply)
		}
	}

	return r
}
