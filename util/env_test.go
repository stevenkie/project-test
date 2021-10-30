package util

import "testing"

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "positive test case",
			want: "local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
