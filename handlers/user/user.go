package user

import (
	"time"
)

type User struct {
	ID                     string    `json:"id" format:"uuid"`
	FirstName              *string   `json:"first_name" nullable:"true"`
	LastName               *string   `json:"last_name" nullable:"true"`
	Email                  *string   `json:"email" nullable:"true" format:"email"`
	PhoneNumber            *string   `json:"phone_number" nullable:"true"`
	PhoneNumberCountryCode *string   `json:"phone_number_country_code" nullable:"true"`
	CreatedAt              time.Time `json:"created_at" format:"date-time" example:"2019-08-24T14:15:22Z"`
}
