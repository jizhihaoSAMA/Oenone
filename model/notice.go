package model

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Notice struct {
    ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
    IsRead      bool               `bson:"is_read" json:"is_read,omitempty"`
    Content     string             `bson:"content" json:"content,omitempty"`
    CreatedTime primitive.DateTime `bson:"created_time" json:"created_time,omitempty"`
}

func NewNotice(content string) *Notice {
    return &Notice{
        ID:          primitive.NewObjectID(),
        CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
        IsRead:      false,
        Content:     content,
    }
}

func NewHouseApprovedNotice(houseID string) *Notice {
    return &Notice{
        ID:          primitive.NewObjectID(),
        CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
        Content:     "<span>恭喜，您的<a href='/index/house/detail/" + houseID + "'>房屋</a>已经审核通过！</span>",
    }
}

func NewHouseRejectedNotice(houseID string) *Notice {
    return &Notice{
        ID:          primitive.NewObjectID(),
        CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
        Content:     "<span>抱歉，您的<a href='/index/house/detail/" + houseID + "'>房屋</a>审核未通过！</span>",
    }
}
