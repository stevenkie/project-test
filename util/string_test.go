package util

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	type args struct {
		e []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{
				e: []string{"a", "b", "b"},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
