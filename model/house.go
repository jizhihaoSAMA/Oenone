package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Online        = 1
	Offline       = 2
	PendingReView = 3
	RejectReview  = 4
)

type House struct {
	ID          string             `json:"id"`
	OwnerID     primitive.ObjectID `json:"owner_id,omitempty"`
	Status      int                `json:"status"`
	CreatedTime time.Time          `json:"created_time"`

	HouseName          string `json:"house_name,omitempty" form:"houseName"`
	HouseLocArea       string `json:"house_loc_area,omitempty" form:"houseLocArea"`
	HouseLocBC         string `json:"house_loc_bc,omitempty" form:"houseLocBc"`
	HouseDetailAddress string `json:"house_detail_address,omitempty" form:"houseDetailAddress"`
	HouseDescription   string `json:"house_description,omitempty" form:"houseDescription"`
	HousePrice         int    `json:"house_price,omitempty" form:"housePrice"`
	StartRentedTime    string `json:"start_rented_time,omitempty" form:"startRentedTime"`
	HouseNeighborhood  string `json:"house_neighborhood,omitempty" form:"houseNeighborhood"`
	HouseSize          int    `json:"house_size,omitempty" form:"houseSize"`
	HouseDirection     string `json:"house_direction,omitempty" form:"houseDirection"`
	FurnitureInfo      int    `json:"furniture_info,omitempty" form:"furnitureInfo"`
	FloorInfo          string `json:"floor_info,omitempty" form:"floorInfo"`
	StructureAmount    string `json:"structure_amount,omitempty" form:"structureAmount"`
	ImageAmount        int    `json:"image_amount"`

	IsFullRent           bool `json:"is_full_rent" form:"isFullRent"`
	SupportShortTermRent bool `json:"support_short_term_rent" form:"supportShortTermRent"`
	HasLift              bool `json:"has_lift" form:"hasLift"`
	HasSingleToilet      bool `json:"has_single_toilet" form:"hasSingleToilet"`
	HasSingleBalcony     bool `json:"has_single_balcony" form:"hasSingleBalcony"`
}

type HouseSearchHit struct {
	ID                string `json:"id,omitempty"`
	HouseName         string `json:"name,omitempty"`
	HouseLocArea      string `json:"house_loc_area,omitempty"`
	HouseLocBC        string `json:"house_loc_bc,omitempty"`
	HouseNeighborhood string `json:"house_neighborhood,omitempty"`
	Content           string `json:"content,omitempty"`
}

type HouseTrend struct {
	ID                string `json:"id,omitempty"`
	HouseName         string `json:"name,omitempty"`
	HouseLocArea      string `json:"house_loc_area,omitempty"`
	HouseLocBC        string `json:"house_loc_bc,omitempty"`
	HouseNeighborhood string `json:"house_neighborhood,omitempty"`

	Score float64 `json:"score,omitempty"`
}

func NewHouseTrend(house *House, score float64) *HouseTrend {
	return &HouseTrend{ID: house.ID, HouseName: house.HouseName, HouseLocArea: house.HouseLocArea, HouseLocBC: house.HouseLocBC, HouseNeighborhood: house.HouseNeighborhood, Score: score}
}
