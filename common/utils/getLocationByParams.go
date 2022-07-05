package utils

import (
	"log"
	"regexp"
	"strconv"
)

type Location struct {
	TopRightLat        string
	TopRightLatFloat   float64
	TopRightLng        string
	TopRightLngFloat   float64
	BottomLeftLat      string
	BottomLeftLatFloat float64
	BottomLeftLng      string
	BottomLeftLngFloat float64
}

func NewLocation(topRightLat string, topRightLng string, bottomLeftLat string, bottomLeftLng string) *Location {
	topRightLatFloat, err := strconv.ParseFloat(topRightLat, 64)
	if err != nil {
		log.Println("[NewLocation] 转换出错，err: ", err)
		return nil
	}
	topRightLngFloat, err := strconv.ParseFloat(topRightLng, 64)
	if err != nil {
		log.Println("[NewLocation] 转换出错，err: ", err)
		return nil
	}
	bottomLeftLatFloat, err := strconv.ParseFloat(bottomLeftLat, 64)
	if err != nil {
		log.Println("[NewLocation] 转换出错，err: ", err)
		return nil
	}
	bottomLeftLngFloat, err := strconv.ParseFloat(bottomLeftLng, 64)
	if err != nil {
		log.Println("[NewLocation] 转换出错，err: ", err)
		return nil
	}
	return &Location{
		TopRightLat:        topRightLat,
		TopRightLng:        topRightLng,
		BottomLeftLat:      bottomLeftLat,
		BottomLeftLng:      bottomLeftLng,
		TopRightLatFloat:   topRightLatFloat,
		TopRightLngFloat:   topRightLngFloat,
		BottomLeftLatFloat: bottomLeftLatFloat,
		BottomLeftLngFloat: bottomLeftLngFloat,
	}
}

func GetLocationByParams(location string) *Location {
	re := regexp.MustCompile("(.*),(.*)-(.*),(.*)")
	result := re.FindStringSubmatch(location)
	if len(result) != 5 {
		log.Println("[GetLocationByParams] 结果不匹配，结果如下：", result)
		return nil
	}

	return NewLocation(result[1], result[2], result[3], result[4])
}
