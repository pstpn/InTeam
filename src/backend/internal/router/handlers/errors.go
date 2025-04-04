package handlers

import (
	"net/http"

	api "backend/internal/router/ogen"
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
	ErrNotModified = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusNotModified,
		Response: api.ErrorResponse{
			Error: api.ErrorsNotModifiedError,
		},
	}
)
