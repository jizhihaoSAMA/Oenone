package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Review struct {
	ID        string          `bson:"_id" json:"id"` // 对应houseID
	Histories []ReviewHistory `bson:"histories" json:"histories"`
}

type ReviewHistory struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	IsApproved bool               `bson:"is_approved" json:"is_approved"`
	ReviewTime primitive.DateTime `bson:"review_time" json:"review_time"`
	ReviewMsg  string             `bson:"review_msg" json:"review_msg"`
}

func NewReview(id string) *Review {
	return &Review{
		ID:        id,
		Histories: []ReviewHistory{},
	}
}

func NewReviewHistory(isApproved bool, reviewMsg string) *ReviewHistory {
	return &ReviewHistory{
		ID:         primitive.NewObjectID(),
		ReviewTime: primitive.NewDateTimeFromTime(time.Now()),
		IsApproved: isApproved,
		ReviewMsg:  reviewMsg,
	}
}
