/* sieves2.go */

package main

import "fmt"

func main() {
	primes := sieve()
	for i := 0; i < 100; i++ {
		fmt.Println(<-primes)
	}
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			out <- prime
			ch = filter(ch, prime)
		}
	}()
	return out
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
