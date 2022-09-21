package validation

import (
	"github.com/modhanami/dinkdonk/user-service/server"
	"github.com/modhanami/dinkdonk/user-service/validation/field"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCreateUserRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *server.CreateUserRequest
		want field.ErrorList
	}{
		{
			name: "empty req",
			req:  &server.CreateUserRequest{},
			want: field.ErrorList{
				field.Required("name", "name is required"),
				field.Required("email", "email is required"),
				field.Required("password", "password is required"),
			},
		},
		{
			name: "invalid email",
			req: &server.CreateUserRequest{
				Name:     "John Doe",
				Email:    "invalid",
				Password: "password",
			},
			want: field.ErrorList{
				field.Invalid("email", "invalid", "email is invalid"),
			},
		},
		{
			name: "too short password",
			req: &server.CreateUserRequest{
				Name:     "John Doe",
				Email:    "john@doe.com",
				Password: "pass",
			},
			want: field.ErrorList{
				field.InvalidLength("password", "pass", 8, 32),
			},
		},
		{
			name: "too long password",
			req: &server.CreateUserRequest{
				Name:     "John Doe",
				Email:    "john@doe.com",
				Password: "passwordpasswordpasswordpassword1",
			},
			want: field.ErrorList{
				field.InvalidLength("password", "passwordpasswordpasswordpassword1", 8, 32),
			},
		},
		{
			name: "valid req",
			req: &server.CreateUserRequest{
				Name:     "John Doe",
				Email:    "john@doe.com",
				Password: "password",
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateCreateUserRequest(tc.req)
			if tc.want == nil {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
