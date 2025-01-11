package user

import (
	"time"
)

type User struct {
	ID                     string  `faker:"uuid"`
	FirstName              *string `faker:"firstName"`
	LastName               *string `faker:"lastName"`
	Email                  *string `faker:"email"`
	PhoneNumber            *string `faker:"phoneNumber"`
	PhoneNumberCountryCode *string `faker:"countryAlpha2"`
	CreatedAt              time.Time
}
