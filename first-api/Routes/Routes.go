package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/user-api")
	{
		grpl.GET("user", Controllers.GetUsers)
		grpl.POST("user", Controllers.CreateUser)
		grpl.GET("user/:id", Controllers.GetUserById)
		grpl.PUT("user/:id", Controllers.UpdateUser)
		grpl.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}
