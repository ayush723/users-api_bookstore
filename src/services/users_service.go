package services

import (
	"github.com/ayush723/users-api_bookstore/utils/crypto_utils"
	"github.com/ayush723/users-api_bookstore/utils/date_utils"
	"github.com/ayush723/utils-go_bookstore/rest_errors"

	"github.com/ayush723/users-api_bookstore/domain/users"
)

var (
	//UsersService is a global variable of type usersServiceInterface equals to the usersService struct
	//we need a variable of type interface to call the methods and this is that variable
	UsersService usersServiceInterface = &usersService{}
)

//usersService struct implements all methods on users.
type usersService struct {
}
//usersServiceInterface includes all methods that can be called only on users
//its basically groupig methods for a specific datatype
type usersServiceInterface interface {
	//GetUser takes a user id from parameter and returns a user
	GetUser(int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest)(*users.User, rest_errors.RestErr)
		}

//GetUser takes a user id from parameter and returns a user
func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	//we never return both. only one at time
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current, err := UsersService.GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if isPartial {

		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s * usersService) LoginUser(request users.LoginRequest)(*users.User, rest_errors.RestErr){
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err :=  dao.FindByEmailAndPassword(); err != nil{
		return nil, err
	}
	return dao, nil
}