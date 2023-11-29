package rest_errors

import "net/http"

type RestErrors struct {
	Message string   `json:"message"`
	Errors  string   `json:"errors"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErrors) Error() string {
	return r.Message
}

func NewRestError(message string, code int64, errors string, causes []Causes) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  errors,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Errors:  "forbidden",
		Code:    http.StatusForbidden,
	}
}
