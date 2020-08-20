package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hojabri/bookstore_users-api/domain/users"
	"github.com/hojabri/bookstore_users-api/services"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me")
}
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
