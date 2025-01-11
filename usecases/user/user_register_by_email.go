package user

import (
	"context"

	"github.com/pioz/faker"
)

type RegisterByEmailRequest struct {
}

type RegisterByEmailResponse struct {
	User User
}

func (u *userUseCaseImpl) UserRegisterByEmail(ctx context.Context, request RegisterByEmailRequest) (*RegisterByEmailResponse, error) {
	mUser := User{}
	err := faker.Build(&mUser)
	if err != nil {
		return nil, err
	}
	return &RegisterByEmailResponse{
		User: mUser,
	}, nil
}
