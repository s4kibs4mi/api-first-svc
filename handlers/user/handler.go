package user

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/s4kibs4mi/api-first-svc/handlers"
	"github.com/s4kibs4mi/api-first-svc/usecases/user"
)

type Handler interface {
	handlers.Handler
	HandleUserRegister(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error)
}

func NewHandler(humaApi huma.API, userUsecase user.UseCase) Handler {
	return &handlerImpl{humaApi: humaApi, userUsecase: userUsecase}
}

type handlerImpl struct {
	humaApi     huma.API
	userUsecase user.UseCase
}

func (h *handlerImpl) Register() {
	h.registerUserRegister()
}
