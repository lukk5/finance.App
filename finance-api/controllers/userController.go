package controllers

import (
	"finance-api-v1/core/middleware"
	"finance-api-v1/core/middleware/handlers"
	"github.com/gin-gonic/gin"
	"time"
)

func UserController(r *gin.Engine, handler handlers.UserHandler) {

	controller := r.Group("v1/api")
	timeout := time.Duration(100) // TODO config use
	controller.Use(middleware.RequestMiddlewareID(), middleware.TimeoutMiddleware(timeout))

	controller.Use()
	{
		//controller.GET("user/get", handler.GetUser(r.c))
	}

	//v1 := r.Group("v1/api")
	//timeout := time.Duration(100) // TODO config use
	//v1.Use(middleware.RequestMiddlewareID(), middleware.TimeoutMiddleware(timeout))

	//v1.Use() // anonymous
	//{
	//	v1.POST("users/login", auth.LoginHandler)
	//}

	//v1.Use(auth.MiddlewareFunc()) // need auth
	//{
	//	v1.GET("users/me", entities.GetUser)
	//}
}
