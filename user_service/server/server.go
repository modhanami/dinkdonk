package server

import (
	"github.com/gin-gonic/gin"
	"github.com/modhanami/dinkdonk/user-service/domain"
	"github.com/modhanami/dinkdonk/user-service/logger"
	"github.com/modhanami/dinkdonk/user-service/validation/field"
)

type UserService interface {
	CreateUser(request *CreateUserRequest) (*domain.User, field.ErrorList, error)
	ListUsers() ([]*domain.User, error)
}

type userHandler struct {
	service UserService
	logger  logger.Logger
}

func NewUserHandler(service UserService, logger logger.Logger) *userHandler {
	return &userHandler{service, logger}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	invalidRequestResponse := NewErrorResponse("Invalid request")
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind json")
		c.JSON(400, invalidRequestResponse)
		return
	}

	_, errList, err := h.service.CreateUser(&req)
	if err != nil {
		h.logger.Error("failed to create user")
		c.JSON(500, NewErrorResponse("failed to create user"))
		return
	}
	if errList != nil {
		c.JSON(400, invalidRequestResponse.WithError(errList))
		return
	}

	c.Status(201)
}

func (h *userHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		h.logger.Error("failed to list users", "error", err)
		c.JSON(500, NewErrorResponse("Failed to list users"))
		return
	}
	c.JSON(200, users)
}

func (h *userHandler) Mount(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.ListUsers)
}
