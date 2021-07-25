package services

import (
	"github.com/ayush723/users-api_bookstore/domain/users"
)

func CreateUser(user users.User) (*users.User, error){
	return &user, nil
}