package crud

import (
	"Oenone/common/base"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"testing"
)

func TestESSearchWithBody(t *testing.T) {
	base.InitService()

	type args struct {
		body gin.H
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{body: gin.H{
				"query": gin.H{
					"match": gin.H{
						"name": "民安",
					},
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ESSearchWithBody(tt.args.body)
			log.Println(got["hits"].)
			if (err != nil) != tt.wantErr {
				t.Errorf("ESSearchWithBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ESSearchWithBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}
