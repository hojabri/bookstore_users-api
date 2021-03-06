package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hojabri/bookstore_users-api/domain/users"
	"github.com/hojabri/bookstore_users-api/services"
	"github.com/hojabri/bookstore_users-api/utils/errors"
	"net/http"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me")
}
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

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}
