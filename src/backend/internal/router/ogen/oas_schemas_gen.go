// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorResponseStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type AuthLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *AuthLoginReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *AuthLoginReq) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *AuthLoginReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *AuthLoginReq) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	Type Types `json:"type"`
}

// GetType returns the value of Type.
func (s *ErrorResponse) GetType() Types {
	return s.Type
}

// SetType sets the value of Type.
func (s *ErrorResponse) SetType(val Types) {
	s.Type = val
}

// ErrorResponseStatusCode wraps ErrorResponse with StatusCode.
type ErrorResponseStatusCode struct {
	StatusCode int
	Response   ErrorResponse
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorResponseStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorResponseStatusCode) GetResponse() ErrorResponse {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorResponseStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorResponseStatusCode) SetResponse(val ErrorResponse) {
	s.Response = val
}

// API Responses.
// Ref: #/components/schemas/LoginResponse
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *LoginResponse) GetAccessToken() string {
	return s.AccessToken
}

// SetAccessToken sets the value of AccessToken.
func (s *LoginResponse) SetAccessToken(val string) {
	s.AccessToken = val
}

// Ref: #/components/schemas/Types
type Types string

const (
	TypesBadRequest       Types = "BadRequest"
	TypesValidationError  Types = "ValidationError"
	TypesNotFoundError    Types = "NotFoundError"
	TypesInternalError    Types = "InternalError"
	TypesNotModifiedError Types = "NotModifiedError"
	TypesConflictError    Types = "ConflictError"
)

// AllValues returns all Types values.
func (Types) AllValues() []Types {
	return []Types{
		TypesBadRequest,
		TypesValidationError,
		TypesNotFoundError,
		TypesInternalError,
		TypesNotModifiedError,
		TypesConflictError,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Types) MarshalText() ([]byte, error) {
	switch s {
	case TypesBadRequest:
		return []byte(s), nil
	case TypesValidationError:
		return []byte(s), nil
	case TypesNotFoundError:
		return []byte(s), nil
	case TypesInternalError:
		return []byte(s), nil
	case TypesNotModifiedError:
		return []byte(s), nil
	case TypesConflictError:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Types) UnmarshalText(data []byte) error {
	switch Types(data) {
	case TypesBadRequest:
		*s = TypesBadRequest
		return nil
	case TypesValidationError:
		*s = TypesValidationError
		return nil
	case TypesNotFoundError:
		*s = TypesNotFoundError
		return nil
	case TypesInternalError:
		*s = TypesInternalError
		return nil
	case TypesNotModifiedError:
		*s = TypesNotModifiedError
		return nil
	case TypesConflictError:
		*s = TypesConflictError
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
