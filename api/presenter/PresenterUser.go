package presenter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"clean_architecture_gin/pkg/entities"

)

type User struct {
	ID          uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Bod         time.Time `json:"bod"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Password    string    `json:"password,omitempty"`
}

func UserSuccessResponse(data *entities.User) *gin.H {
	user := User{
		ID:          data.ID,
		Name:        data.Name,
		Username:    data.Username,
		Gender:      data.Gender,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Bod:         data.Bod,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Password:    data.Password,
	}

	return &gin.H{
		"status": true,
		"result": user,
	}
}

func UsersSuccessResponse(data []User) *gin.H {
	var users []User
	for _, item := range data {
		users = append(users, User{
			ID:          item.ID,
			Name:        item.Name,
			Username:    item.Username,
			Gender:      item.Gender,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			Bod:         item.Bod,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return &gin.H{
		"status": true,
		"result": users,
	}
}

func UserErrorResponse(err error) *gin.H {
	return &gin.H{
		"status": false,
		"errors": err.Error(),
	}
}
