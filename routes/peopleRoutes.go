package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/go-crud/controllers"
	"github.com/yosefalemu/go-crud/initializers"
)

func RegisterPeopleRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/people", func(c *gin.Context) {
			controllers.GetPerson(c, initializers.DB)
		})
		v1.POST("/people", func(c *gin.Context) {
			controllers.CreatePerson(c, initializers.DB)
		})
	}
}
