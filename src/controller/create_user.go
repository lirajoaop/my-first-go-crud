package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lirajoaop/my-first-go-crud/src/configuration/validation"
	"github.com/lirajoaop/my-first-go-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
