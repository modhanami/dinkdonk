package gormrepository

import (
	"github.com/modhanami/dinkdonk/user-service/domain"
	"github.com/modhanami/dinkdonk/user-service/logger"
	"github.com/modhanami/dinkdonk/user-service/model/gormmodel"
	"github.com/modhanami/dinkdonk/user-service/service"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

var _ service.UserRepository = (*userRepository)(nil)

func NewUserRepository(db *gorm.DB, logger logger.Logger) service.UserRepository {
	return &userRepository{db, logger}
}

func (s *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	userModel := &gormmodel.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := s.db.Create(userModel).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &domain.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}, nil
}

func (s *userRepository) ListUsers() ([]*domain.User, error) {
	var users []*gormmodel.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	var usersDomain []*domain.User
	for _, user := range users {
		usersDomain = append(usersDomain, &domain.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	return usersDomain, nil
}
