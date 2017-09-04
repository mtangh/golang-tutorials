/* exec.go */

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("dscl", ".", "-list", "/Users")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	if output := stdout.String(); len(output) > 0 {
		fmt.Print(output)
	}
	if output := stderr.String(); len(output) > 0 {
		fmt.Print(output)
	}

	os.Exit(0)
}
