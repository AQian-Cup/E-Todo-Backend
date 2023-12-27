package errno

import "e-todo-backend/pkg/response"

var (
	UserAlreadyExist = &response.Response{
		HTTP: 400,
		Result: response.ErrorResult{
			Code:    "FailedOperation.UserAlreadyExist",
			Message: "User already exist",
		},
	}
	UserAuthError = &response.Response{
		HTTP: 401,
		Result: response.ErrorResult{
			Code:    "FailedOperation.UserAuthError",
			Message: "User authentication error",
		},
	}
)
