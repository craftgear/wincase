package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func wincase(filename string) string {
	var cases = []struct {
		input  string
		output string
	}{
		{"<", "＜"},
		{">", "＞"},
		{":", "："},
		{"\"", "”"},
		{"/", "／"},
		{"\\", "＼"},
		{"|", "｜"},
		{"?", "？"},
		{"*", "＊"},
	}
	for _, v := range cases {
		if strings.Index(filename, v.input) > -1 {
			filename = strings.Replace(filename, v.input, v.output, -1)
		}
	}
	return strings.Trim(filename, " ")
}

func ren(old, new string, dryRun, verbose bool) error {
	if dryRun {
		fmt.Printf("rename %s to %s\n", color.RedString(old), color.GreenString(new))
		return nil
	} else {
		if verbose {
			fmt.Printf("renaming: %s to %s\n", color.RedString(old), color.GreenString(new))
		}
		return os.Rename(old, new)
	}
}

//TODO traverseという関数を作って、その中でWalkする
//func traverse() error {

//}

func main() {
	showHelp := func() {
		fmt.Println("\n ＊ wincase - make files live even on windows\n")
		fmt.Println("\twincase is a simple utility to recursively replace\n\tforbidden characters on Windows platforms\n\twith 2-byte corresponding characters\n")
		fmt.Println("Usage: wincase [options] target_dir\n")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var dir string
	var help bool
	var dryRun bool
	var verbose bool

	//コマンドラインオプション解析
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&dryRun, "-dry-run", false, "dry run")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	if help {
		showHelp()
	}

	// 対象ディレクトリ
	if args := flag.Args(); len(args) > 0 {
		dir = args[0]
	} else {
		showHelp()
		//_dir, err := os.Getwd()
		//if err != nil {
		//log.Fatal(err)
		//}
		//dir = _dir
	}

	dirStack := []string{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dir {
			return nil
		}

		filename := filepath.Base(path)
		parentDir := filepath.Dir(path)

		if strings.Index(filename, ".") == 0 {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if casedName := wincase(filename); filename != casedName {
			if info.IsDir() {
				dirStack = append(dirStack, path)
			} else {
				newPath := filepath.Join(parentDir, casedName)
				err := ren(path, newPath, dryRun, verbose)
				if err != nil {
					fmt.Printf("err = %+v\n", err)
				}
			}

		}
		return nil
	})

	for i := len(dirStack) - 1; i >= 0; i-- {
		path := dirStack[i]

		casedName := wincase(filepath.Base(path))
		parentDir := filepath.Dir(path)

		newPath := filepath.Join(parentDir, casedName)
		err := ren(path, newPath, dryRun, verbose)
		if err != nil {
			fmt.Printf("err = %+v\n", err)
		}
	}

	if verbose {
		fmt.Println("done.")
	}
}

// TODO マルチプラットフォームバイナリ生成
// TODO ドキュメントにインストール方法と使い方を書く
