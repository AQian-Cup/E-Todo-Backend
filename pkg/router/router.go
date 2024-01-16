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
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware(userController.PublicKey))
		{
			users.GET("/current", userController.ReadCurrent)
		}
		taskController := &controllerV1.TaskController{}
		tasks := v1.Group("/tasks")
		tasks.Use(middleware.AuthMiddleware(userController.PublicKey))
		{
			tasks.POST("/", taskController.Create)
			tasks.DELETE("/", taskController.Delete)
			tasks.DELETE("/:id", taskController.Delete)
			tasks.GET("/", taskController.ReadList)
			tasks.GET("/:id", taskController.Read)
			tasks.PATCH("/:id", taskController.Edit)
		}
	}
	return nil
}
