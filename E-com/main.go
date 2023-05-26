package main

import (
	"E-com/controllers"
	"log"
	"os"

	"E-com/middleware"

	"github.com/gin-gonic/gin"

  "E-com/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	// other routes

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

  log.Fatal(router.Run(":" + port));
}
