package errno

import "e-todo-backend/pkg/response"

var (
	UserAlreadyExist = &response.StructErrorResponse{
		HTTP: 400,
		Result: response.Result{
			Code:    "FailedOperation.UserAlreadyExist",
			Message: "User already exist",
		},
	}
)
