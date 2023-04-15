package routes

import (
	"JWT/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incomingRoutes *gin.Engine){
	 incomingRoutes.POST("user/signup", controllers.Signup());
	 incomingRoutes.POST("user/login", controllers.Login());
}