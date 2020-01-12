package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"os"
)

var log = golog.New("main")

const (
	Version = "2.9.2"
)

func main() {

	flag.Parse()

	// 版本
	if *paramVersion {
		fmt.Printf("%s", Version)
		return
	}

	switch *paramMode {
	case "v3":
		V3Entry()
	default:
		fmt.Println("--mode not specify")
		os.Exit(1)
	}

}
