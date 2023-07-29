package handlers

import (
	"github.com/gin-gonic/gin"

	"clean_architecture_gin/api/presenter"
	"clean_architecture_gin/pkg/entities"
	"clean_architecture_gin/pkg/user"

)

func IndexHandler(service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := service.IndexService()
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		c.JSON(200, presenter.UsersSuccessResponse(results))
	}
}

func StoreHandler(service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body entities.User

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		result, err := service.StoreService(&body)
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		c.JSON(201, presenter.UserSuccessResponse(result))
	}
}

func ShowHandler(service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := service.ShowService(id)
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		c.JSON(200, presenter.UserSuccessResponse(result))
	}
}

func UpdateHandler(service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var body entities.User
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		result, err := service.UpdateService(id, &body)
		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		c.JSON(200, presenter.UserSuccessResponse(result))
	}
}

func DeleteHandler(service user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := service.DeleteService(id)

		if err != nil {
			c.JSON(500, presenter.UserErrorResponse(err))
		}

		c.JSON(200, presenter.UserDeleteResponse(result))
	}
}
