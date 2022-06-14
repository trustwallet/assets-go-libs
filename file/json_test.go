package file

import (
	"reflect"
	"testing"
)

func TestPrepareJSONData(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				payload: struct {
					Name string `json:"name"`
				}{Name: "\u003ctest"},
			},
			wantErr: false,
			want:    []byte("{\n    \"name\": \"<test\"\n}"),
		},
		{
			name: "wrong_json",
			args: args{
				payload: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrepareJSONData(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareJSONData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareJSONData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
