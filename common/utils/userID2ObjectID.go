package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func UserID2ObjectID(userID string) primitive.ObjectID {
	o, _ := primitive.ObjectIDFromHex(userID)
	return o
}
