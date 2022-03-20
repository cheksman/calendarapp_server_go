package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Fullname     string             `json:"fullname" bson:"fullname"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Description  string             `json:"description" bson:"description"`
	ProfileImage string             `json:"profileImage" bson:"profileImage"`
	Address      Address            `json:"address" bson:"address"`
}
