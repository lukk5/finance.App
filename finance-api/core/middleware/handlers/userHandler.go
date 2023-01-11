package handlers

import (
	"context"
	"finance-api-v1/core/middleware/mainhandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	GetUser(ctx context.Context)
}

type Handler struct {
	testas
}

func GetUser(context *gin.Context) {
	mainhandler.HandleRequest(context, func(c *gin.Context) *mainhandler.Response {
		currentUser := User{ID: 1, UserName: "testas"}

		return mainhandler.NewSuccessResponse(http.StatusOK, currentUser)
	})
}

func NewUserHandler() *Handler {
	return &UserHandler{}
}
