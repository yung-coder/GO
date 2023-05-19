package main

import (
	"E-com/controllers"
	"E-com/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"E-com/middleware"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("waiting...")
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
}
