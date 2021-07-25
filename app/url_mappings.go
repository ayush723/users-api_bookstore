package app

import (
	"github.com/ayush723/users-api_bookstore/controllers/ping"
	"github.com/ayush723/users-api_bookstore/controllers/users"
)
func mapUrls(){
	router.GET("/ping", ping.Ping)
	
	router.GET("/users/:user-id", users.GetUser)
	router.GET("/users/search",users.SearchUser)
	router.POST("/users", users.CreateUser)
}