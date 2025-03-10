package handlers

import (
	api "backend/internal/router/ogen"
	"net/http"
)

var (
	ErrInternal = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.ErrorResponse{
			Error: api.ErrorsInternalError,
		},
	}
	ErrBadRequest = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response: api.ErrorResponse{
			Error: api.ErrorsBadRequest,
		},
	}
	ErrUnauthorized = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusUnauthorized,
		Response: api.ErrorResponse{
			Error: api.ErrorsUnauthorized,
		},
	}
)
