/* user.go */

package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Printf("user: ERROR: %s\n", err.Error())
	}

	groupIds, err := u.GroupIds()

	fmt.Printf("user %T %v\n", u, u)
	fmt.Printf("user.Name = %s\n", u.Name)
	fmt.Printf("user.Username = %v\n", u.Username)
	fmt.Printf("user.Uid = %v\n", u.Uid)
	fmt.Printf("user.Gid = %v\n", u.Gid)
	fmt.Printf("user.GroupIds = %v\n", groupIds)
	fmt.Printf("user.HomeDir = %s\n", u.HomeDir)

	return
}
