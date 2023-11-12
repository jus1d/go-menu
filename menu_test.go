package menu

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		options []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Menu
		wantErr bool
	}{
		{
			name:    "should error",
			args:    args{options: []string{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{options: []string{"option 1"}},
			want: &Menu{
				options: []string{"option 1"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
