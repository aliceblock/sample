package routes

import (
	"net/http"
	"os"

	"github.com/aliceblock/sample/internal/mode"

	"github.com/aliceblock/sample/cmd/server/repository"
	"github.com/gin-gonic/gin"
)

// UserRoutes generate '/users' for Gin RouterGroup
func UserRoutes(r *gin.RouterGroup) {
	dataMode := os.Getenv("DATAMODE")
	var m mode.DataMode
	if dataMode == "release" {
		m = mode.PROD
	} else {
		m = mode.MOCK
	}
	userRepo := repository.NewUserRepo(m)
	user := r.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": userRepo.Model.GetData(),
			})
		})
	}
}
