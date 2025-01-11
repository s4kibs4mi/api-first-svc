package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/api-first-svc/handlers"
	"github.com/s4kibs4mi/api-first-svc/usecases/user"
)

const userRegisterDescription = `
This endpoint allows users to register for the platform. Users can register using one of the following methods:
- **Email and Password**: Users provide their email and password to create an account.
- **Phone Number**: Users provide their phone number to register.

Additionally, users can optionally provide their first name and last name during registration.


**Note**: Either email and password or phone must be provided; providing both is optional but not required.
If both email and phone number are provided, the account will be created using the email.
`

func (h *handlerImpl) registerUserRegister() {
	huma.Register[RegisterRequest, RegisterResponse](
		h.humaApi,
		huma.Operation{
			OperationID:   "user_register",
			Method:        http.MethodPost,
			Path:          handlers.BuildPathWithV1("users/register"),
			Summary:       "Register an user",
			Description:   userRegisterDescription,
			Tags:          []string{"USERS"},
			DefaultStatus: http.StatusCreated,
			Responses: map[string]*huma.Response{
				"201": {
					Content: map[string]*huma.MediaType{
						"application/json": {
							Schema: &huma.Schema{
								Ref: "#/components/schemas/RegisterResponseData",
							},
						},
					},
				},
				"400": {
					Content: map[string]*huma.MediaType{
						"application/json": {
							Schema: &huma.Schema{
								Ref: "#/components/schemas/ErrorModel",
							},
						},
					},
				},
				"500": {
					Content: map[string]*huma.MediaType{
						"application/json": {
							Schema: &huma.Schema{
								Ref: "#/components/schemas/ErrorModel",
							},
						},
					},
				},
			},
		},
		h.HandleUserRegister,
	)
}

type RegisterRequest struct {
	Body struct {
		FirstName                *string    `json:"first_name" required:"false" nullable:"true" minLength:"1"`
		LastName                 *string    `json:"last_name" required:"false" nullable:"true" minLength:"1"`
		Email                    *string    `json:"email" required:"false" nullable:"true" format:"email" dependentRequired:"password"`
		Password                 *string    `json:"password" required:"false" nullable:"true" minLength:"8"`
		PhoneNumber              *string    `json:"phone_number" required:"false" nullable:"true" minLength:"10" format:"phone" dependentRequired:"phone_number_country_code_id"`
		PhoneNumberCountryCodeID *uuid.UUID `json:"phone_number_country_code_id" required:"false" nullable:"true" format:"uuid"`
	} `contentType:"application/json"`
}

func (i RegisterRequest) Resolve(ctx huma.Context, prefix *huma.PathBuffer) []error {
	if i.Body.Email == nil && i.Body.PhoneNumber == nil {
		return []error{
			&huma.ErrorDetail{
				Location: "body.email",
				Message:  "is required",
			},
			&huma.ErrorDetail{
				Location: "body.password",
				Message:  "is required",
			},
		}
	}
	return nil
}

type RegisterResponseData struct {
	Data *User `json:"data,omitempty" required:"false"`
}

type RegisterResponse struct {
	Body RegisterResponseData `contentType:"application/json"`
}

func (h *handlerImpl) HandleUserRegister(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	if request.Body.Email != nil {
		resp, err := h.userUsecase.UserRegisterByEmail(ctx, user.RegisterByEmailRequest{})
		if err != nil {
			return nil, &huma.ErrorDetail{
				Message: err.Error(),
			}
		}
		return &RegisterResponse{
			Body: RegisterResponseData{
				Data: &User{
					ID:                     resp.User.ID,
					FirstName:              resp.User.FirstName,
					LastName:               resp.User.LastName,
					Email:                  resp.User.Email,
					PhoneNumber:            resp.User.PhoneNumber,
					PhoneNumberCountryCode: resp.User.PhoneNumberCountryCode,
					CreatedAt:              resp.User.CreatedAt,
				},
			},
		}, nil
	}

	resp, err := h.userUsecase.UserRegisterByPhone(ctx, user.RegisterByPhoneRequest{})
	if err != nil {
		return nil, &huma.ErrorDetail{
			Message: err.Error(),
		}
	}
	return &RegisterResponse{
		Body: RegisterResponseData{
			Data: &User{
				ID:                     resp.User.ID,
				FirstName:              resp.User.FirstName,
				LastName:               resp.User.LastName,
				Email:                  resp.User.Email,
				PhoneNumber:            resp.User.PhoneNumber,
				PhoneNumberCountryCode: resp.User.PhoneNumberCountryCode,
				CreatedAt:              resp.User.CreatedAt,
			},
		},
	}, nil
}
