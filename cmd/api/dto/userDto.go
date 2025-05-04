package dto

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Id, validation.Required, is.UUID),
	)
}
