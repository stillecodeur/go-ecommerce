package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"id"`
	FirstName      string             `json:"firstName" bson:"first_name" validate:"required,min=2,max=30"`
	LastName       string             `json:"lastName" bson:"last_name" validate:"required,min=2,max=30"`
	Password       string             `json:"password" bson:"password" validate:"required,min=6"`
	Email          string             `json:"email" bson:"email" validate:"email,required"`
	Phone          string             `json:"phone" bson:"phone" validate:"required"`
	Token          string             `json:"token" bson:"token"`
	RefreshToken   string             `json:"refreshToken" bson:"refresh_token"`
	CreatedAt      time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updated_at"`
	UserID         string             `json:"userId" bson:"user_id"`
	UserCart       []ProductUser      `json:"userCart" bson:"user_cart"`
	AddressDetails []Address          `json:"addressDetails" bson:"address_details"`
	OrderStatus    []Order            `json:"orderStatus" bson:"order_status"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"id"`
	ProductName *string            `json:"productName" bson:"product_name"`
	Price       *uint64            `json:"price" bson:"price"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"id"`
	ProductName *string            `json:"productName" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Rating      *uint              `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"id"`
	House     *string            `json:"house" bson:"house"`
	Street    *string            `json:"street" bson:"street"`
	City      *string            `json:"city" bson:"city"`
	Pincode   *string            `json:"pinCode" bson:"pincode"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"id"`
	OrderCart     []ProductUser      `json:"orderCart" bson:"order_cart"`
	OrderedAt     time.Time          `json:"orderedAt" bson:"ordered_at"`
	Price         int                `json:"price" bson:"price"`
	Discount      *int               `json:"dicsount" bson:"discount"`
	PaymentMethod Payment            `json:"paymentMethod" bson:"payment_method"`
}

type Payment struct {
	Digial bool
	COD    bool
}
