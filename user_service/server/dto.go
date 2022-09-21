package server

import "strings"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateUserRequest) Sanitize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)
}

type CreateUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

func (e *ErrorResponse) WithError(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: e.Message,
		Error:   err,
	}
}
