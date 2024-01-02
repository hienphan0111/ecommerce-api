package database

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't decode products")
	ErrUserIdIsNotValid   = errors.New("User id is not valid")
	ErrCantUpdateUser     = errors.New("Can't update user")
	ErrCantRemoveItemCart = errors.New("Can't remove item from cart")
	ErrCantGetItemCart    = errors.New("Can't get item from cart")
	ErrCantBuyFromCart    = errors.New("Can't buy from cart")
)

func AddProductToCart() {

}

func RemoveItemFromCart() {

}

func GetItemFromCart() {

}

func BuyFromCart() {

}

func InstantBuy() {

}
