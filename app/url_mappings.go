package app

import (
	"github.com/ayush723/users-api_bookstore/controllers/ping"
	"github.com/ayush723/users-api_bookstore/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user-id", users.Get)
	router.PUT("/users/:user-id", users.Update)
	router.PATCH("/users/:user-id", users.Update)
	router.DELETE("/users/:user-id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
