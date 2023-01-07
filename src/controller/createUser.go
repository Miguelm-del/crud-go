package controller

import (
	"fmt"
	"github.com/Miguelm-del/crud-go/src/config/rest_err"
	"github.com/Miguelm-del/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadResquestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
