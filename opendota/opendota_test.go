package opendota

import (
	"testing"
)

func TestGetMatchPlayers(t *testing.T) {
	type args struct {
		playerID string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "TB哥", args: args{"147805190"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			players := GetMatchPlayers(tt.args.playerID)
			t.Logf("%+v", players)
		})
	}
}
