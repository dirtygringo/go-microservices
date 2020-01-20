package users

import (
	"strings"

	"github.com/dirtygringo/go-microservices/bookstore_users_api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	if user.Id <= 0 {
		return errors.BadRequest("invalid user id")
	}
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	return nil
}
