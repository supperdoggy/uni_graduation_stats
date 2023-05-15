package utils

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/db"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 50

// ValidateUserEmailAndPassword validates the email and password, returns error if invalid
func ValidateUserEmailAndPassword(email, password string) error {
	if err := validation.Validate(&email, validation.Required, is.Email); err != nil {
		return errors.New("invalid email")
	}

	// validate password
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		return errors.New("password is too simple")
	}
	return nil
}

// MapDBUserToResponseUser takes an input of type db.User and
// returns a pointer to rest.User struct.
// The new instance has ID, Email, CreatedAt, and EditedAt fields set to corresponding values of the input
func MapDBUserToResponseUser(u db.User) *rest.User {
	return &rest.User{
		ID:       u.ID,
		Email:    u.Email,
		FullName: u.FullName,

		CreatedAt: u.CreatedAt,
		EditedAt:  u.EditedAt,
	}
}
