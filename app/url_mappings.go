package app

import (
	"github.com/ayush723/users-api_bookstore/controllers/ping"
	"github.com/ayush723/users-api_bookstore/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user-id", users.GetUser)
	router.PUT("/users/:user-id", users.UpdateUser)
	router.PATCH("/users/:user-id", users.UpdateUser)

	router.GET("/users/search", users.SearchUser)
}
