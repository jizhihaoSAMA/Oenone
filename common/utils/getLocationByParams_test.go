package utils

import (
	"log"
	"reflect"
	"testing"
)

func TestGetLocationByParams(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name string
		args args
		want *Location
	}{
		{
			"1",
			args{location: "1.9,2-2,3"},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocationByParams(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocationByParams() = %v, want %v", got, tt.want)
			} else {
				log.Println(got)
			}
		})
	}
}
