package routes

import (
	"db/controller"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incomingRoutes *gin.Engine){
	 incomingRoutes.POST("user/signup", controller.Signup());
	 incomingRoutes.POST("user/login", controller.Login());
}