/* server.go */

package main

import "fmt"

type request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *request) {
	reply := op(req.a, req.b)
	fmt.Printf("run: %T(%v) -> reply %v\n", req, req, reply)
	req.replyc <- reply
}

func server(op binOp, service chan *request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (service chan *request, quit chan bool) {
	service = make(chan *request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	for i := N - 1; i >= 0; i-- {
		if answer := <-reqs[i].replyc; answer != N+2*i {
			fmt.Println("fail at", i)
		}
	}
	fmt.Println("done")
	quit <- true
}
