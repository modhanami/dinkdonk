package service

import (
	"github.com/modhanami/dinkdonk/user-service/domain"
	"github.com/modhanami/dinkdonk/user-service/server"
	"github.com/modhanami/dinkdonk/user-service/validation"
	"github.com/modhanami/dinkdonk/user-service/validation/field"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}

type userService struct {
	repo UserRepository
}

func (u *userService) ListUsers() ([]*domain.User, error) {
	return u.repo.ListUsers()
}

func NewUserService(repo UserRepository) server.UserService {
	return &userService{repo}
}

var _ server.UserService = (*userService)(nil)

func (u *userService) CreateUser(req *server.CreateUserRequest) (*domain.User, field.ErrorList, error) {
	req.Sanitize()
	errList := validation.ValidateCreateUserRequest(req)
	if errList != nil {
		return nil, errList, nil
	}
	user, err := u.repo.CreateUser(&domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, nil, err
	}
	return user, nil, nil
}
