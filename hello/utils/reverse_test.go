package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	arr := []struct {
		in, want string
	}{
		{"Hello world", "dlrow olleH"},
		{"你好世界", "界世好你"},
		{"", ""},
	}

	for _, c := range arr {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
