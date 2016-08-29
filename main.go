package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var dir string
	var help bool

	//コマンドラインオプション解析
	flag.BoolVar(&help, "h", false, "show help")
	flag.Parse()

	if help {
		fmt.Println("\nwincase - a simple utility to replace forbidden characters on Windows platofrms with 2-byte corresponding characters\n")
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

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("path = %+v\n", path)
		return err
	})
}
