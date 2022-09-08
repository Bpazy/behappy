package kda

import "testing"

func TestGetKda(t *testing.T) {
	type args struct {
		kills  int
		deaths int
		assist int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				kills:  4,
				deaths: 3,
				assist: 15,
			},
			want: "6.3",
		},
		{
			args: args{
				kills:  5,
				deaths: 0,
				assist: 5,
			},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run("KDA", func(t *testing.T) {
			if got := GetKda(tt.args.kills, tt.args.deaths, tt.args.assist); got != tt.want {
				t.Errorf("GetKda() = %v, want %v", got, tt.want)
			}
		})
	}
}
