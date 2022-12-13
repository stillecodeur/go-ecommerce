package database

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this is user is not valid")
	ErrCantUpdateUser     = errors.New("can't update this user")
	ErrCantRemoveItemCart = errors.New("can't remove item from cart")
	ErrCantGetItem        = errors.New("can't get item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddToCart() gin.HandlerFunc {
	database
}

func RemoveItem() gin.HandlerFunc {

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
