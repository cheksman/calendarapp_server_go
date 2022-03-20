package models

type User struct {
	Fullname     string     `json:"fullname" bson:"fullname"`
	Email        string     `json:"email" bson:"email"`
	Password     string     `json:"password" bson:"password"`
	Description  string     `json:"description" bson:"description"`
	ProfileImage string     `json:"profileImage" bson:"profileImage"`
	Address      Address    `json:"address" bson:"address"`
	Calendar     []Calendar `json:"calendar" bson:"calendar"`
}
