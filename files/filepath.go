/* filepath.go */

package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user.HomeDir = %s\n", u.HomeDir)

	dir := filepath.Join(u.HomeDir, ".config", "myapp")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return
}
