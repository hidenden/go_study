/*
 - コマンドライン引数の取り扱い｡ os.Args []string
 - getopt相当の処理  package flag
 - 環境変数の扱い （標準・サードパーティPKG)
 - コンフィグレーションファイル的なサポート
*/

package main

import (
	"fmt"
	"os"

	"github.com/hidenden/go_study/argenv/option"
)

func main() {
	//	args()
	if err := options(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}

/*
func args() {
	fmt.Printf("Length of os.Args:%d\n", len(os.Args))
	fmt.Println("-- Dump args --")
	for i, v := range os.Args {
		fmt.Printf("[%d]:%s\n", i, v)
	}
}
*/

func options() error {
	if err := option.Parse(); err != nil {
		return err
	}

	fmt.Printf("Log Directory: %s\n", option.Dir())
	fmt.Printf("Port number: %d\n", option.Port())
	fmt.Printf("Combined mode: %v\n", option.Combined())
	fmt.Printf("Subdir mode: %v\n", option.Subdir())

	return nil
}
