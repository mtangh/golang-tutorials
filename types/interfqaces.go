/* interfaces.go */

package main

import "fmt"

type Any interface{}

func main() {

	var obj Any

	obj = 12
	fmt.Println("obj =", obj)

	if v, ok := obj.(int); ok {
		fmt.Println("obj is int;", v)
	}

	obj = "abc"
	fmt.Println("obj =", obj)

	if v, ok := obj.(string); ok {
		fmt.Println("obj is string;", v)
	}

	str := "defg"
	obj = &str
	fmt.Println("obj =", obj)

	if v, ok := obj.(*string); ok {
		fmt.Println("obj is *string;", *v)
		*v = "hijk"
		fmt.Println("obj is *string;", str)
	}

	var any Any = "defg"
	obj = &any
	fmt.Println("obj =", obj)

	if v, ok := obj.(*Any); ok {
		fmt.Println("obj is *Any;", v, any)
	}

}
