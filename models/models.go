package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	First_Name    *string            `json:"first_name"`
	Last_Name     *string            `json:"last_name"`
	Email         *string            `json:"email"`
	Password      *string            `json:"password"`
	Phone_Number  *string            `json:"phone_number"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	User_ID       string             `json:"user_id"`
	User_Cart     []ProductUser      `json:"user_cart"`
	Address       []Address          `json:"address"`
	Order_Status  []Order            `json:"order_status"`
}

type Product struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name        *string            `json:"product_name"`
	Product_Description *string            `json:"product_description"`
	Price               *float64           `json:"price"`
	Rating              *float64           `json:"rating"`
	Image               *string            `json:"image"`
}

type ProductUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name string             `json:"product_name"`
	Price        *float64           `json:"price"`
	Rating       *float64           `json:"rating"`
	Image        *string            `json:"image"`
}

type Address struct {
	Address_ID string  `json:"address_id"`
	House      *string `json:"house"`
	Street     *string `json:"street"`
	City       *string `json:"city"`
	Pincode    *string `json:"pincode"`
}

type Order struct {
	Order_ID       string        `json:"order_id"`
	Order_Cart     []ProductUser `json:"order_cart"`
	Ordered_At     time.Time     `json:"ordered_at"`
	Price          float64       `json:"price"`
	Discount       *float64      `json:"discount"`
	Payment_Method Payment       `json:"payment_method"`
}

type Payment struct {
	Payment_ID     string   `json:"payment_id"`
	Payment_Method *string  `json:"payment_method"`
	Amount         *float64 `json:"amount"`
	COD            *bool    `json:"cod"`
}
