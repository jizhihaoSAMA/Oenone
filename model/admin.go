package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username,omitempty"`
	Password string             `bson:"password" json:"-"`
	Role     int                `bson:"role" json:"role,omitempty"`
}

func NewAdmin(username string, password string, role int) *Admin {
	return &Admin{ID: primitive.NewObjectID(), Username: username, Password: password, Role: role}
}
