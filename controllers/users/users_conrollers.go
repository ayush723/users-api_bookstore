package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ayush723/users-api_bookstore/domain/users"
	"github.com/ayush723/users-api_bookstore/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err!=nil{
		restErr := errors.RestErr{
			Message:"invalid json body",
			Code:"http.StatusBadRequest",
			Error:"bad_request",
		}
		fmt.Println(err)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr!=nil{
		fmt.Println(saveErr)
		return
	}
	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context){

		c.String(http.StatusNotImplemented, "implement me")
}
func SearchUser(c *gin.Context){}

