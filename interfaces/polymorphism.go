/* polymorphism.go */

package main

import "fmt"

type Object interface {
}

type ToString interface {
	String() string
}

type TestType struct {
	a int
	b string
}

type TestType2 struct {
	a int
	b string
}

func (t *TestType) String() string {
	return "a=" + fmt.Sprintf(t.b) + " b=" + t.b
}

func main() {
	var t1 Object
	var t2 Object
	t1 = &TestType{77, "Sunset Strip"}
	t2 = &TestType2{77, "Sunset Strip"}
	fmt.Println("[t1]")
	if r, ok := t1.(ToString); ok {
		fmt.Println("t1: r is ToString")
		fmt.Printf("t1: r=%T\n", r)
		fmt.Print("t1: ")
		fmt.Println(r)
	} else {
		fmt.Println("t1: r is not ToString")
		fmt.Printf("t1: %v\n", t1)
	}
	fmt.Println("[t2]")
	if r, ok := t2.(ToString); ok {
		fmt.Println("t2: t2 is ToString")
		fmt.Printf("t2: r=%T\n", r)
		fmt.Print("t2: ")
		fmt.Println(r)
	} else {
		fmt.Println("t2: t2 is not ToString")
		fmt.Printf("t2: %v\n", t2)
	}
	return
}
