package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/middleware"
)


func  UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate());
	incomingRoutes.Use("/users", controller.GetUsers());
	incomingRoutes.Use("/users/:user_id" , controller.GetUser());
}