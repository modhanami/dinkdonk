package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "empty email",
			arg:  "",
			want: false,
		},
		{
			name: "invalid email",
			arg:  "invalid",
			want: false,
		},
		{
			name: "valid email",
			arg:  "john@doe.com",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidEmail(tc.arg)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "empty password",
			arg:  "",
			want: false,
		},
		{
			name: "too short password",
			arg:  "pass",
			want: false,
		},
		{
			name: "too long password",
			arg:  "passwordpasswordpasswordpassword1",
			want: false,
		},
		{
			name: "valid password",
			arg:  "password",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidPassword(tc.arg)
			assert.Equal(t, tc.want, got)
		})
	}
}
