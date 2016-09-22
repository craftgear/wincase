package main

import (
	"os"
	"testing"

	"github.com/craftgear/fileutils"
)

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
	src := "./test"
	dest := "./test2"

	err := ren(src, dest, false, false)
	if err != nil {
		t.Error(err)
	}
	t.Log(fileutils.Exist(dest))
	if !fileutils.Exist(dest) {
		t.Errorf("%s should exist", dest)
	}

	//dry run
	err = ren(dest, src, true, false)
	if err != nil {
		t.Error(err)
	}

	if fileutils.Exist(src) {
		t.Errorf("%s should not exist", src)
	}

	err = ren(dest, src, false, false)
	if err != nil {
		t.Error(err)
	}
}

func TestTraverse(t *testing.T) {
	//testディレクトリを_testとしてコピーしてバックアップする
	testDir := "./test"
	testDirBackup := "./_test"

	err := fileutils.CpDir(testDirBackup, testDir, 0755)
	if err != nil {
		t.Error(err)
	}

	//testディレクトリに対してdryRun testディレクトリ以下はそのまま
	traverse(testDir, true, false)
	if fileutils.Exist(testDir+"/hoge*") == false {
		t.Error("renaming should've not done.")
	}

	//testディレクトリに対してリネーム実行 testディレクトリいかがリネームされている
	traverse(testDir, false, false)
	if fileutils.Exist(testDir + "/hoge*") {
		t.Error("renaming should've done.")
	}
	//testディレクトリを削除
	os.RemoveAll(testDir)

	//_testディレクトリをtestディレクトリにリネーム
	os.Rename(testDirBackup, testDir)
}
