package todo

import (
	"gin-money-manager-api/modules/todo/query"
	"gin-money-manager-api/modules/todo/repository"
	"gin-money-manager-api/modules/todo/resource"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(
	r *gin.Engine,
	db *gorm.DB,
) {
	todoRepository := repository.NewTodoRepository(db)

	todoResource := resource.NewTodoResource(todoRepository)

	TodoFindCompleted := query.NewTodoFindCompleted(todoRepository)

	todos := r.Group("/todos")
	{
		todos.GET("/completed", TodoFindCompleted.Handler)

		todos.GET("", todoResource.Index)
		todos.GET("/:id", todoResource.Show)
		todos.POST("", todoResource.Create)
		todos.DELETE("/:id", todoResource.Delete)
		todos.PATCH("/:id", todoResource.Update)
	}
}
