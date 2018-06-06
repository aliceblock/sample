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
			html := "<h1>Root /API</h1>"

			//Write your 200 header status (or other status codes, but only WriteHeader once)
			c.Writer.WriteHeader(http.StatusOK)
			//Convert your cached html string to byte array
			c.Writer.Write([]byte(html))
		})
	}
	UserRoutes(api)
}
