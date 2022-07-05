package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Appointment struct {
	HouseID       string             `bson:"_id" json:"house_id,omitempty"`
	CreatedTime   primitive.DateTime `bson:"created_time" json:"created_time,omitempty"`
	AppointedTime primitive.DateTime `bson:"appointed_time" json:"appointed_time,omitempty"`
}

type AppointmentDto struct {
	*Appointment
	HouseName         string `json:"house_name,omitempty"`
	HouseLocArea      string `json:"house_loc_area,omitempty"`
	HouseLocBC        string `json:"house_loc_bc,omitempty"`
	HouseNeighborhood string `json:"house_neighborhood,omitempty"`
}

func NewAppointmentDto(appointment *Appointment, house *House) *AppointmentDto {
	return &AppointmentDto{Appointment: appointment, HouseName: house.HouseName, HouseLocArea: house.HouseLocArea, HouseLocBC: house.HouseLocBC, HouseNeighborhood: house.HouseNeighborhood}
}

func NewAppointment(houseID string, appointedTime primitive.DateTime) *Appointment {
	return &Appointment{
		HouseID:       houseID,
		CreatedTime:   primitive.NewDateTimeFromTime(time.Now()),
		AppointedTime: appointedTime,
	}
}
