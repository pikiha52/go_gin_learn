package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"clean_architecture_gin/api/presenter"
	"clean_architecture_gin/pkg/entities"

)

type Repository interface {
	IndexRepository() ([]presenter.User, error)
	StoreRepository(user *entities.User) (*entities.User, error)
	ShowRepository(id string) (*entities.User, error)
	UpdateRepository(id string, user *entities.User) (*entities.User, error)
	DeleteRepository(id string) (*entities.User, error)
}

type repository struct {
	Database *gorm.DB
}

func NewRepo(database *gorm.DB) Repository {
	return &repository{
		Database: database,
	}
}

func (r *repository) IndexRepository() ([]presenter.User, error) {
	var users []presenter.User
	r.Database.Find(&users)

	return users, nil
}

func (r *repository) StoreRepository(user *entities.User) (*entities.User, error) {
	err := r.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) ShowRepository(id string) (*entities.User, error) {
	var user *entities.User
	r.Database.Where("id = ?", id).First(&user)

	return user, nil
}

func (r *repository) UpdateRepository(id string, user *entities.User) (*entities.User, error) {
	var userModel entities.User
	r.Database.Where("id = ?", id).First(&userModel)
	if userModel.ID == uuid.Nil {
		return nil, errors.New("User not found!")
	}

	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Gender = user.Gender
	userModel.Address = user.Address
	userModel.PhoneNumber = user.PhoneNumber
	userModel.Bod = user.Bod
	userModel.Password = user.Password
	userModel.DeletedAT = user.DeletedAT

	r.Database.Save(&userModel)

	return &userModel, nil
}

func (r *repository) DeleteRepository(id string) (*entities.User, error) {
	var user *entities.User

	r.Database.Where("id = ?", id).First(&user)

	err := r.Database.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
