package middleware

import (
	"log"
	"net/http"

	"blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
				utils.ErrorWithStatus(c, http.StatusInternalServerError, 500, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
