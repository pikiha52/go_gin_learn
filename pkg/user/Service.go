package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"clean_architecture_gin/api/presenter"
	"clean_architecture_gin/pkg/entities"
)

type Service interface {
	IndexService() ([]presenter.User, error)
	StoreService(user *entities.User) (*entities.User, error)
	ShowService(id string) (*entities.User, error)
	UpdateService(id string, user *entities.User) (*entities.User, error)
	DeleteService(id string) (*entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) IndexService() ([]presenter.User, error) {
	return s.repository.IndexRepository()
}

func (s *service) StoreService(user *entities.User) (*entities.User, error) {
	user.ID = uuid.New()

	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	return s.repository.StoreRepository(user)
}

func (s *service) ShowService(id string) (*entities.User, error) {
	data, err := s.repository.ShowRepository(id)
	if err != nil {
		return nil, err
	}

	if data.ID == uuid.Nil {
		return nil, errors.New("User not found!")
	}

	return s.repository.ShowRepository(id)
}

func (s *service) UpdateService(id string, user *entities.User) (*entities.User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	return s.repository.UpdateRepository(id, user)
}

func (s *service) DeleteService(id string) (*entities.User, error) {
	data, err := s.repository.DeleteRepository(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
