package service

import (
	"github.com/dirtygringo/go-microservices/bookstore_users_api/domain/users"
	"github.com/dirtygringo/go-microservices/bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.BadRequest("invalid user id")
	}
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
