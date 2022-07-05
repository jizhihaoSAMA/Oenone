package public

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"reflect"
	"testing"
)

func TestGetGroupByPlaces1(t *testing.T) {
	base.InitService()
	type args struct {
		field  string
		places []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"1",
			args{
				field:  utils.ToSnakeCase("houseLocArea"),
				places: []string{"西城区", "东城区"},
			}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGroupByPlaces(tt.args.field, tt.args.places); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupByPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
