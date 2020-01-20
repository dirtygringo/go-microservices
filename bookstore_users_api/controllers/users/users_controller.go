package users

import (
	"net/http"
	"strconv"

	"github.com/dirtygringo/go-microservices/bookstore_users_api/domain/users"
	"github.com/dirtygringo/go-microservices/bookstore_users_api/service"
	"github.com/dirtygringo/go-microservices/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := service.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("invalid user id type")
		c.JSON(err.Status, err)
		return
	}
	result, err := service.GetUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
