package option

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var dir string

const (
	expDir = "Log directory [required]"
)

func Dir() string {
	return dir
}

var port uint

const (
	defaultPort = 8080
	expPort     = "Port number"
)

func Port() uint {
	return port
}

var combined bool

const (
	defaultCombined = false
	expCombined     = "Save log data with combined mode"
)

func Combined() bool {
	return combined
}

var subdir bool

const (
	defaultSubdir = false
	expSubdir     = "Save each logfiles under dated sub directory"
)

func Subdir() bool {
	return subdir
}

func init() {
	flag.StringVar(&dir, "d", "", expDir)
	flag.StringVar(&dir, "dir", "", expDir)
	flag.UintVar(&port, "p", defaultPort, expPort)
	flag.UintVar(&port, "port", defaultPort, expPort)
	flag.BoolVar(&combined, "c", defaultCombined, expCombined)
	flag.BoolVar(&combined, "combined", defaultCombined, expCombined)
	flag.BoolVar(&subdir, "s", defaultSubdir, expSubdir)
	flag.BoolVar(&subdir, "subdir", defaultSubdir, expSubdir)
}

//  初期化はinitで行う｡他に必要なI/Fは
//  - オプションのパースを実施する
//     指定されたディレクトリの正当性も確認する｡
//  - オプションの値を取り出す
//		取り出すためのアクセサをパッケージとして提供する｡

func Parse() error {
	flag.Parse() // Default is ExitOnError

	if dir == "" {
		return errors.New("log directory is not specified")
	}
	if err := checkdir(dir); err != nil {
		return err
	}
	return nil
}

func checkdir(d string) error {
	fi, err := os.Stat(d)
	if err != nil {
		return fmt.Errorf("log directory check error:%w", err)
	}
	if !fi.IsDir() {
		return errors.New(d + " is not a directory")
	}
	return nil
}
