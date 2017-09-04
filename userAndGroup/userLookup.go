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

		if userInfo, err := user.Lookup(flag.Arg(index)); err == nil {
			fmt.Printf("%s; %s:%s:%s:*:%s:%s:*\n",
				flag.Arg(index),
				userInfo.Username, userInfo.Uid, userInfo.Gid,
				userInfo.Name, userInfo.HomeDir)
		}

	}

	os.Exit(0)
}
