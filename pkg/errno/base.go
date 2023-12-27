package errno

import "e-todo-backend/pkg/response"

var (
	OK = &response.Response{
		HTTP:   200,
		Result: response.OkResult{},
	}
	InternalServerError = &response.Response{
		HTTP: 500,
		Result: response.ErrorResult{
			Code:    "InternalError",
			Message: "Internal server error",
		},
	}
	PageNotFound = &response.Response{
		HTTP: 404,
		Result: response.ErrorResult{
			Code:    "ResourceNotFound.PageNotFound",
			Message: "Page not found",
		},
	}
	BindError = &response.Response{
		HTTP: 400,
		Result: response.ErrorResult{
			Code:    "InvalidParameter.BindError",
			Message: "Error occurred while binding the request body to the struct",
		},
	}
	InvalidParameter = &response.Response{
		HTTP: 400,
		Result: response.ErrorResult{
			Code:    "InvalidParameter",
			Message: "Parameter verification failed.",
		},
	}
)
