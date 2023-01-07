package routes

import (
	"github.com/Miguelm-del/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userID", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userID", controller.UpdateUserByID)
	r.DELETE("/deleteUser/:userID", controller.DeleteUser)

}
