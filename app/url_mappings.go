package app

import (
	"github.com/abhishekgbhat-yb/bookstore_users_api/controllers/ping"
	"github.com/abhishekgbhat-yb/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)

}
