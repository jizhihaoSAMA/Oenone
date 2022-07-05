package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Star struct {
	HouseID     string             `bson:"_id" json:"house_id,omitempty"`
	CreatedTime primitive.DateTime `bson:"created_time" json:"created_time,omitempty"`
}

type StarDto struct {
	*Star
	HouseName         string `json:"house_name,omitempty"`
	HouseLocArea      string `json:"house_loc_area,omitempty"`
	HouseLocBC        string `json:"house_loc_bc,omitempty"`
	HouseNeighborhood string `json:"house_neighborhood,omitempty"`
}

func NewStarDto(house *House, star *Star) *StarDto {
	return &StarDto{
		Star:              star,
		HouseName:         house.HouseName,
		HouseLocArea:      house.HouseLocArea,
		HouseLocBC:        house.HouseLocBC,
		HouseNeighborhood: house.HouseNeighborhood,
	}
}

func NewStar(houseID string) *Star {
	return &Star{
		CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
		HouseID:     houseID,
	}
}
