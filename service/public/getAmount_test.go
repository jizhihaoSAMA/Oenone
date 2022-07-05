package public

import (
	"Oenone/common/base"
	"log"
	"testing"
)

func TestGetAmountOfAllHouse(t *testing.T) {
	base.InitService()
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			"1",
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAmountOfHouse()
			log.Println(got)
			if got != tt.want {
				t.Errorf("GetAmountOfHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}
