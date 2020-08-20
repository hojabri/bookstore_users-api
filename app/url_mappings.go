package app

import (
	"github.com/hojabri/bookstore_users-api/controllers/ping"
	"github.com/hojabri/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search",controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
