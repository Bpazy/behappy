package bjson

import "testing"

type temp struct {
	Content string `json:"content"`
}

func TestMustJsonMarshal(t *testing.T) {

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Normal", args: args{temp{"This is content"}}, want: "{\"content\":\"This is content\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustJsonMarshal(tt.args.v); got != tt.want {
				t.Errorf("MustJsonMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustJsonUnmarshal(t *testing.T) {
	type args struct {
		b   []byte
		got interface{}
	}
	tests := []struct {
		name string
		args args
		want temp
	}{
		{name: "Normal", args: args{b: []byte("{\"content\":\"This is content\"}"), got: &temp{}}, want: temp{"This is content"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustJsonUnmarshal(tt.args.b, tt.args.got)
			if *tt.args.got.(*temp) != tt.want {
				t.Errorf("MustJsonUnmarshal().got = %v, want %v", tt.args.got, tt.want)
			}
		})
	}
}
