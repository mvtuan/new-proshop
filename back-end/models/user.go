package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"last_name,omitempty"`
	IsAdmin   bool               `json:"isAdmin,omitempty" bson:"is_admin,omitempty"`
}
