package user

import "context"

type UseCase interface {
	UserRegisterByEmail(ctx context.Context, request RegisterByEmailRequest) (*RegisterByEmailResponse, error)
	UserRegisterByPhone(ctx context.Context, request RegisterByPhoneRequest) (*RegisterByPhoneResponse, error)
}

type userUseCaseImpl struct {
}

func NewUseCase() UseCase {
	return &userUseCaseImpl{}
}
