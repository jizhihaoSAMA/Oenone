package public

import (
	"Oenone/common/base"
	"reflect"
	"testing"
)

func TestSearchNeighborhood(t *testing.T) {
	base.InitService()

	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{{
		args: args{
			"民安",
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchNeighborhood(tt.args.query, "东城区", 3)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchNeighborhood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchNeighborhood() got = %v, want %v", got, tt.want)
			}
		})
	}
}
