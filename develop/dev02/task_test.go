package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"", ""}, {"asd", "asd"}, {"a2e", "aae"},
		{`a4bc2d5e`, "aaaabccddddde"}, {`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"}, {`qwe\\5`, `qwe\\\\\`}, {"kk", "kk"},
		{"66", ""}, {"", ""},
	}
	for _, tt := range tests {
		testname := "test1"
		t.Run(testname, func(t *testing.T) {
			ans, _ := unpack(tt.a)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
