package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIRoutes generate '/api' for Gin Engine Router
func APIRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "/web/templates/api.tmpl", gin.H{})
		})
	}
	UserRoutes(api)
}
