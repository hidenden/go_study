package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s %s\n", get_hello(), get_world())
}

func get_hello() string {
	return "Hello"
}

func get_world() string {
	return "World"
}
