package validation

import (
	"github.com/modhanami/dinkdonk/user-service/server"
	"github.com/modhanami/dinkdonk/user-service/validation/field"
)

func ValidateCreateUserRequest(user *server.CreateUserRequest) field.ErrorList {
	var errs field.ErrorList
	if user.Name == "" {
		errs = append(errs, field.Required("name", "name is required"))
	}
	if user.Email == "" {
		errs = append(errs, field.Required("email", "email is required"))
	}
	if user.Password == "" {
		errs = append(errs, field.Required("password", "password is required"))
	}
	if len(errs) > 0 {
		return errs
	}

	if !IsValidEmail(user.Email) {
		errs = append(errs, field.Invalid("email", user.Email, "email is invalid"))
	}
	if !IsValidPassword(user.Password) {
		errs = append(errs, field.InvalidLength("password", user.Password, 8, 32))
	}
	if len(errs) > 0 {
		return errs
	}

	return nil
}
