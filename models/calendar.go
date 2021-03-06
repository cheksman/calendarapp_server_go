package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Country string `json:"country" bson:"country"`
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Street  string `json:"street" bson:"street"`
	Zipcode string `json:"zipcode" bson:"zipcode"`
}

type Calendar struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	Date        string             `json:"date" bson:"date"`
	Start       string             `json:"start" bson:"start"`
	End         string             `json:"end" bson:"end"`
	Description string             `json:"description" bson:"description"`
	Link        string             `json:"link" bson:"link"`
	Address     Address            `json:"address" bson:"address"`
}
