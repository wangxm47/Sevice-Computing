package say_hello

import "testing"

func TestSay_hello(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"wyx", "wyx welcome to Golang World!!!"},
		{"wyk", "wyk welcome to Golang World!!!"},
		{"", " welcome to Golang World!!!"},
	}
	for _, c := range cases {
		got := Say_hello(c.in)
		if got != c.want {
			t.Errorf("Say_hello(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
