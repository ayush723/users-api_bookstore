package users

import (
	"fmt"

	"github.com/ayush723/users-api_bookstore/utils/errors"
)

//only has access to database
var(
	usersDB = make(map[int64]*User)
)

//dao = data access object

func (user *User) Get()*errors.RestErr{
	result := usersDB[user.Id]
	if result == nil{
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found",user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr{
	current := usersDB[user.Id]
	
	if  current != nil{
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists",user.Id))
	}
	usersDB[user.Id] = user
	return nil
}