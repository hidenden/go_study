package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Now %v\n", now)
	fmt.Printf("ANSIC %s\n", now.Format(time.ANSIC))
	fmt.Printf("RFC3339 %s\n", now.Format(time.RFC3339))
	fmt.Printf("ISO8601 %s\n", now.Format("2006-01-02"))

	dir_name := now.Format("2006-01-02")
	fmt.Printf("Temp dir name:%s\n", dir_name)

}
