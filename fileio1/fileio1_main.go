package main

import (
	"fmt"
	"os"

	"github.com/hidenden/go_study/fileio1/gocat"
)

func showhow() {
	fmt.Fprintf(os.Stderr, "gocat <sub command> <file name>\n")
	fmt.Fprintf(os.Stderr, "  sub command: line|bytes|string|scanner\n")
	os.Exit(1)
}

func getCat(cmd string, f string) gocat.Cat {
	var c gocat.Cat
	o := os.Stdout

	switch cmd {
	case "line":
		c = gocat.NewCatReadLine(f, o)
	case "bytes":
		c = gocat.NewCatReadBytes(f, o)
	case "string":
		c = gocat.NewCatReadString(f, o)
	case "scan":
		c = gocat.NewCatScanner(f, o)
	default:
		showhow()
	}
	return c
}

func main() {
	if len(os.Args) != 3 {
		showhow()
	}
	c := getCat(os.Args[1], os.Args[2])
	if err := gocat.Read(c); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
