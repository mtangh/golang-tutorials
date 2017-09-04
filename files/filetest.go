/* filetest.go */

package main

import (
	"fmt"
	"os"
)

import (
	"./file"
)

func main() {
	hello := []byte("heelo, world\n")
	file.Stdout.Write(hello)
	f, err := file.Open("/tmp/test")
	if f == nil {
		fmt.Printf("can't open file; err: %s\n", err.Error())
		os.Exit(1)
	}
}
