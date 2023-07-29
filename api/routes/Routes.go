package routes

import (
	"github.com/gin-gonic/gin"

	"clean_architecture_gin/api/handlers"
	"clean_architecture_gin/pkg/user"

)

func Routes(api *gin.RouterGroup, service user.Service) {
	api.GET("/user", handlers.IndexHandler(service))
	api.POST("/user", handlers.StoreHandler(service))
	api.GET("/user/:id", handlers.ShowHandler(service))
	api.PUT("/user/:id", handlers.UpdateHandler(service))
	api.DELETE("/user/:id", handlers.DeleteHandler(service))
}
