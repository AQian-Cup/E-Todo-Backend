package router

import (
	"crypto/ecdsa"
	controllerV1 "e-todo-backend/pkg/controller/v1"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/jwt"
	"e-todo-backend/pkg/middleware"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		response.Write(c, errno.PageNotFound)
	})
	v1 := g.Group("/v1")
	{
		userController := &controllerV1.UserController{
			PrivateKey: nil,
			PublicKey:  nil,
		}
		userController.PrivateKey, _ = jwt.GenerateECPrivateKey()
		userController.PublicKey = userController.PrivateKey.Public().(*ecdsa.PublicKey)
		_ = v1.Group("/users")
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		taskController := &controllerV1.TaskController{}
		task := v1.Group("/tasks")
		task.Use(middleware.AuthMiddleware(userController.PublicKey))
		{
			task.POST("/", taskController.Create)
			task.DELETE("/", taskController.Delete)
			task.DELETE("/:id", taskController.Delete)
			task.GET("/", taskController.ReadList)
			task.GET("/:id", taskController.Read)
			task.PATCH("/:id", taskController.Edit)
		}
	}
	return nil
}
