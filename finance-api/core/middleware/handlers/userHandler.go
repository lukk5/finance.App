package handlers

import (
	userRepo "finance-api-v1/core/database/repo"
	//entity "finance-api-v1/core/entities"
	"finance-api-v1/core/middleware/mainhandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	GetUser(c *gin.Context, id uint)
	//UpdateUser(ctx context.Context, user *entity.User)
}

type Handler struct {
	userRepo userRepo.UserRepo
}

func (h *Handler) GetUser(c *gin.Context, id uint) {

	var user, err = h.userRepo.Get(c, id)

	if err != nil {
		mainhandler.NewInternalErrorResponse(err)
	}

	mainhandler.HandleRequest(c, func(c *gin.Context) *mainhandler.Response {
		return mainhandler.NewSuccessResponse(http.StatusOK, user)
	})
}

func NewUserHandler(repo userRepo.UserRepo) UserHandler {
	return &Handler{
		userRepo: repo,
	}
}
