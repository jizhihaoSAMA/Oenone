package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreatedTime  primitive.DateTime `bson:"created_time" json:"created_time"`
	Username     string             `json:"username" gorm:"varchar(100);not null" bson:"username"`
	Telephone    string             `json:"telephone" gorm:"varchar(100);not null;unique" bson:"telephone"`
	Password     string             `json:"password" gorm:"varchar(100)" bson:"password"`
	UnreadAmount int                `json:"unread_amount" bson:"unread_amount"`
	Stars        []*Star            `bson:"stars" json:"stars"`
	Notices      []*Notice          `bson:"notices" json:"notices"`
	Appointments []*Appointment     `bson:"appointments" json:"appointments"`
}

func NewUser(username string, password string, telephone string) *User {
	return &User{
		ID:          primitive.NewObjectID(),
		CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
		Username:    username,
		Telephone:   telephone,
		Password:    password,
		Stars:       []*Star{},
		Notices:     []*Notice{},
	}
}

type UserDto struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}
