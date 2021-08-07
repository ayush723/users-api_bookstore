package users

import (
	"net/http"
	"strconv"

	"github.com/ayush723/oauth-go_bookstore/oauth"
	"github.com/ayush723/utils-go_bookstore/rest_errors"

	"github.com/ayush723/users-api_bookstore/domain/users"
	"github.com/ayush723/users-api_bookstore/services"
	"github.com/gin-gonic/gin"
)

func TestServiceInterface() {

}

//getUserId converts obtained user id into int64
func getUserId(userIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErrs := strconv.ParseInt(userIdParam, 10, 64)
	if userErrs != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")

	}
	return userId, nil
}
//Create is a handler to create new user
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-public") == "true"))
}

//Get is a handler to get existing user
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil{
		c.JSON(err.Status, err)
		return
	}

	//userid is a int64 user id
	userId, idErr := getUserId(c.Param("user-id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	//GetUser returns a user by its user id
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	//GetCallerId get caller id from header
	if oauth.GetCallerId(c.Request) == user.Id{
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

//Update is a handler to update existing user
func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user-id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-public") == "true"))

}

//Delete is a handler to delete any user
func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user-id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

//Search is a handler to search user 
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-public") == "true"))
}

func Login(c *gin.Context){
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request) ; err != nil{
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err!= nil{
		c.JSON(err.Status, err)
		return 
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-public") == "true"))
	}
