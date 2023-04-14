package controllers

import (
	models "JWT/Models"
	"JWT/database"
	"JWT/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	validate "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.New().Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count1, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "While counting"})
			return
		}

		count2, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "While counting  phone"})
			return
		}

		if count1 > 0 || count2 > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Account already exists"})
		}
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, user)
	}
}
