package errno

import "e-todo-backend/pkg/response"

var (
	OK = &response.StructDataResponse{
		HTTP: 200,
		Result: response.Result{
			Code:    "",
			Message: "",
		},
	}
	InternalServerError = &response.StructErrorResponse{
		HTTP: 500,
		Result: response.Result{
			Code:    "InternalError",
			Message: "Internal server error",
		},
	}
	PageNotFound = &response.StructErrorResponse{
		HTTP: 404,
		Result: response.Result{
			Code:    "ResourceNotFound.PageNotFound",
			Message: "Page not found",
		},
	}
	BindError = &response.StructErrorResponse{
		HTTP: 400,
		Result: response.Result{
			Code:    "InvalidParameter.BindError",
			Message: "Error occurred while binding the request body to the struct",
		},
	}
	InvalidParameter = &response.StructErrorResponse{
		HTTP: 400,
		Result: response.Result{
			Code:    "InvalidParameter",
			Message: "Parameter verification failed.",
		},
	}
)
