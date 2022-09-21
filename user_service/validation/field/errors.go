package field

import "fmt"

type ErrorList []*Error

func (e ErrorList) Error() string {
	return fmt.Sprintf("%+v", "")
}

type Error struct {
	Type     string      `json:"type"`
	Field    string      `json:"field"`
	BadValue interface{} `json:"bad_value,omitempty"`
	Detail   string      `json:"detail"`
}

func (e Error) Error() string {
	if e.BadValue == nil {
		return fmt.Sprintf("%s: %s", e.Field, e.Detail)
	}
	return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Detail, e.BadValue)
}

var _ error = (*Error)(nil)
var _ error = (*ErrorList)(nil)

const (
	ErrorTypeRequired      = "required"
	ErrorTypeInvalid       = "invalid"
	ErrorTypeInvalidLength = "invalid_length"
)

func Required(field string, detail string) *Error {
	return &Error{
		Type:     ErrorTypeRequired,
		Field:    field,
		BadValue: nil,
		Detail:   detail,
	}
}

func Invalid(field string, value interface{}, detail string) *Error {
	return &Error{
		Type:     ErrorTypeInvalid,
		Field:    field,
		BadValue: value,
		Detail:   detail,
	}
}

func InvalidLength(field string, value interface{}, min, max int) *Error {
	return &Error{
		Type:     ErrorTypeInvalidLength,
		Field:    field,
		BadValue: value,
		Detail:   fmt.Sprintf("must be between %d and %d characters", min, max),
	}
}
