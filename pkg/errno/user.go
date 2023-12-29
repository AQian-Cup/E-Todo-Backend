package errno

import (
	"e-todo-backend/pkg/response"
	"net/http"
)

var (
	UserAlreadyExist = &response.Response{
		HTTP: http.StatusBadRequest,
		Result: response.ErrorResult{
			Code:    "FailedOperation.UserAlreadyExist",
			Message: "User already exist",
		},
	}
	UserAuthError = &response.Response{
		HTTP: http.StatusUnauthorized,
		Result: response.ErrorResult{
			Code:    "FailedOperation.UserAuthError",
			Message: "User authentication error",
		},
	}
)
