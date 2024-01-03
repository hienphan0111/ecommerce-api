package controllers

import (
	"log"
	"net/http"

	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hienphan0111/ecommerce-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// HashPassword is a function that hashes the password
func HashPassword(password string) string {

}

// ComparePassword is a function that compares the password
func VerifyPassword(password string, hashedPassword string) (bool, string) {

}

// SignUp is a function that signs up a user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
			return
		}

	}
}

// Login is a function that logs in a user
func Login() gin.HandlerFunc {

}

// ProductViewerAdmin is a function that adds a product to the database
func ProductViewerAdmin() gin.HandlerFunc {

}

// SearchProduct is a function that searches for a product
func SearchProduct() gin.HandlerFunc {

}

// SearchProductByQuery is a function that searches for a product by query
func SearchProductByQuery() gin.HandlerFunc {

}
