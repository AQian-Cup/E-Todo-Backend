package errno

import (
	"e-todo-backend/pkg/response"
	"net/http"
)

var (
	OK = &response.Response{
		HTTP:   http.StatusOK,
		Result: response.OkResult{},
	}
	InternalServerError = &response.Response{
		HTTP: http.StatusInternalServerError,
		Result: response.ErrorResult{
			Code:    "InternalError",
			Message: "Internal server error",
		},
	}
	PageNotFound = &response.Response{
		HTTP: http.StatusNotFound,
		Result: response.ErrorResult{
			Code:    "ResourceNotFound.PageNotFound",
			Message: "Page not found",
		},
	}
	BindError = &response.Response{
		HTTP: http.StatusBadRequest,
		Result: response.ErrorResult{
			Code:    "InvalidParameter.BindError",
			Message: "Error occurred while binding the request body to the struct",
		},
	}
	InvalidParameter = &response.Response{
		HTTP: http.StatusBadRequest,
		Result: response.ErrorResult{
			Code:    "InvalidParameter",
			Message: "Parameter verification failed.",
		},
	}
)
