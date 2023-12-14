package router

import (
	controllerV1 "e-todo-backend/pkg/controller/v1"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		response.Write(c, errno.PageNotFound)
	})
	v1 := g.Group("/v1")
	{
		userController := &controllerV1.UserController{}
		user := v1.Group("/user")
		{
			user.POST("/register", userController.Register)
		}
	}
	return nil
}
