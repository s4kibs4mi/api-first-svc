package user

import (
	"context"

	"github.com/pioz/faker"
)

type RegisterByPhoneRequest struct{}
type RegisterByPhoneResponse struct {
	User User
}

func (u *userUseCaseImpl) UserRegisterByPhone(ctx context.Context, request RegisterByPhoneRequest) (*RegisterByPhoneResponse, error) {
	mUser := User{}
	err := faker.Build(&mUser)
	if err != nil {
		return nil, err
	}
	return &RegisterByPhoneResponse{
		User: mUser,
	}, nil
}
