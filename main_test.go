package main

import "testing"

func TestWinCase(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{in: "foo*bar*", out: "foo＊bar＊"},
		{in: "foo<bar<", out: "foo＜bar＜"},
		{in: "foo>bar>", out: "foo＞bar＞"},
		{in: "foo:bar:", out: "foo：bar："},
		{in: "foo/bar/", out: "foo／bar／"},
		{in: "foo|bar|", out: "foo｜bar｜"},
		{in: "foo?bar?", out: "foo？bar？"},
		{in: "foo\"bar\"", out: "foo”bar”"},
		{in: "foo\\bar\\", out: "foo＼bar＼"},
		{in: "foo_bar_", out: "foo_bar_"},
		{in: "foo_bar ", out: "foo_bar"},
	}

	for _, c := range cases {
		if e := wincase(c.in); e != c.out {
			t.Errorf("%s should be %s but got %s", c.in, c.out, e)
		}
	}
}

func TestRen(t *testing.T) {
	// TODO

}
