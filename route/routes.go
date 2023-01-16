package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, v1 *gin.RouterGroup, db *gorm.DB) {
	routeBody := IntegrateFile(db)

	r.GET("api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!",
		})
	})

	v1.GET("/user", routeBody.UserHandler.GetAllUser)
	v1.GET("/user/:id", routeBody.UserHandler.GetUserById)
	v1.POST("user", routeBody.UserHandler.CreateUser)
	v1.PUT("/user/:id", routeBody.UserHandler.UpdateUser)
	v1.DELETE("/user/:id", routeBody.UserHandler.DeleteUser)

}
