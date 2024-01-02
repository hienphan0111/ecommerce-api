package controllers

import (
	"github.com/gin-gonic/gin"
)

// HashPassword is a function that hashes the password
func HashPassword(password string) string {

}

// ComparePassword is a function that compares the password
func VerifyPassword(password string, hashedPassword string) (bool, string) {

}

// SignUp is a function that signs up a user
func SignUp() gin.HandlerFunc {

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
