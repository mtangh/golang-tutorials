/* userLookup.go */

package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		os.Exit(1)
	}
	for index := 0; index < flag.NArg(); index++ {

		if len(flag.Arg(index)) <= 0 {
			continue
		}

		if group, err := user.LookupGroup(flag.Arg(index)); err == nil {
			fmt.Printf("%s; %s:%s*\n",
				flag.Arg(index),
				group.Name, group.Gid)
		}

	}

	os.Exit(0)
}
