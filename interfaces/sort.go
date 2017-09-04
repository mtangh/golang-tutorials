/* sort.go */

package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func main() {
	ints()
}

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 909, 0, 0, 42, 7586, -5467984, 7586}
	a := IntSlice(data)
	sort.Sort(a)
	if sort.IsSorted(a) {
		fmt.Print("OK\n")
	} else {
		panic("fail")
	}
	return
}

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	return
}
