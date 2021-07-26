package users

import (
	"net/http"
	"strconv"

	"github.com/ayush723/users-api_bookstore/utils/errors"

	"github.com/ayush723/users-api_bookstore/domain/users"
	"github.com/ayush723/users-api_bookstore/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, userErrs := strconv.ParseInt(c.Param("user-id"), 10, 64)
	if userErrs != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func SearchUser(c *gin.Context) {}
