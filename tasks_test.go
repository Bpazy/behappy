package behappy

import (
	"github.com/Bpazy/behappy/dto"
	"reflect"
	"testing"
)

func Test_getNewMatchPlayersByMatchId(t *testing.T) {
	type args struct {
		subNewMatchPlayers []*dto.MatchPlayer
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[int64][]*dto.MatchPlayer
	}{
		{name: "tmp", args: struct{ subNewMatchPlayers []*dto.MatchPlayer }{
			[]*dto.MatchPlayer{
				{MatchID: 1, PlayerID: "1"},
				{MatchID: 1, PlayerID: "2"},
				{MatchID: 2, PlayerID: "2"},
			},
		}, wantResult: map[int64][]*dto.MatchPlayer{
			1: {{MatchID: 1, PlayerID: "1"}, {MatchID: 1, PlayerID: "2"}},
			2: {{MatchID: 2, PlayerID: "2"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getNewMatchPlayersByMatchId(tt.args.subNewMatchPlayers); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("getNewMatchPlayersByMatchId() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
