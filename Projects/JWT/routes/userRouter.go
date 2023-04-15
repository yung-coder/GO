package routes

import (
	"JWT/controllers"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/middleware"
)


func  UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate());
	incomingRoutes.Use("/users", controllers.GetUsers());
	incomingRoutes.Use("/users/:user_id" , controllers.GetUser());
}