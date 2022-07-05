package public

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"encoding/json"
	"log"
	"testing"
)

func TestGetGroupByLocation(t *testing.T) {
	base.InitService()
	type args struct {
		location *utils.Location
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{utils.NewLocation("39.933475", "116.414290", "39.933470", "116.414295")},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := json.Marshal(GetGroupByLocation(tt.args.location))
			log.Println(string(res))
		})
	}
}
