package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentAddReq is type of insert studen response
type StudentAddReq struct {
	FirstName string `json:"first_name,omitempty" validate:"required,min=3"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
	ClassName string `json:"class_name,omitempty" validate:"required"`
}

type StudentUpdateReq struct {
	StudentAddReq
	Id string `json:"id" validate:"required"`
}

type StudentSearchRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
}

// Student struct for model
type Student struct {
	ID         primitive.ObjectID `bson:"_id"`
	SequenceID int                `bson:"id"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	ClassName  string             `bson:"class_name"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type UserAddReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type TokenClaims struct {
	ID    string
	Email string
}
