package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//TODO テスト書く

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
	return filename
}

func main() {
	var dir string
	var help bool
	var dryRun bool

	//コマンドラインオプション解析
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&dryRun, "dry-run", false, "dry run")
	flag.Parse()

	if help {
		fmt.Println("\nwincase - make files live even on windows\n wincase is a simple utility to replace forbidden characters on Windows platofrms with 2-byte corresponding characters\n")
		fmt.Println("Usage: wincase [options] target_dir")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// 対象ディレクトリ
	if args := flag.Args(); len(args) > 0 {
		dir = args[0]
	} else {
		dir = "./"
	}

	dirStack := []string{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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
				if dryRun {
					fmt.Printf("rename %s to %s\n", path, newPath)
				} else {
					err := os.Rename(path, newPath)
					if err != nil {
						return err
					}
				}
			}

		}
		return nil
	})

	//fmt.Printf("dirStack = %+v\n", dirStack)

	for i := len(dirStack) - 1; i >= 0; i-- {
		path := dirStack[i]

		casedName := wincase(filepath.Base(path))
		parentDir := filepath.Dir(path)

		newPath := filepath.Join(parentDir, casedName)
		if dryRun {
			fmt.Printf("rename %s to %s\n", path, newPath)
		} else {
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf("err = %+v\n", err)
			}
		}
	}
}
