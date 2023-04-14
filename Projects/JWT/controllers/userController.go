package controllers

import (
	"JWT/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()
