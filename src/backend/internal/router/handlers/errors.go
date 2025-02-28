package handlers

import (
	api "backend/internal/router/ogen"
	"net/http"
)

var (
	ErrInternal = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.ErrorResponse{
			Type: api.TypesInternalError,
		},
	}
	ErrBadRequest = &api.ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response: api.ErrorResponse{
			Type: api.TypesBadRequest,
		},
	}
)
