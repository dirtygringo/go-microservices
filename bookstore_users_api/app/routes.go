package app

import (
	"github.com/dirtygringo/go-microservices/bookstore_users_api/controllers/ping"
	"github.com/dirtygringo/go-microservices/bookstore_users_api/controllers/users"
)

func mapRoutes() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
