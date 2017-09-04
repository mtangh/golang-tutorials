/* rot13cat.go */

package main

import (
	"flag"
	"fmt"
	"os"
)

import (
	"./catRot13"
	"./file"
)

var rot13Flag = flag.Bool("rot13", false, "rot13 the input")

func cat(r catRot13.Reader) {
	const NBUF = 512
	var buff [NBUF]byte

	if *rot13Flag {
		r = catRot13.NewRotate13(r)
	}
	for {
		switch nr, er := r.Read(buff[:]); true {
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := file.Stdout.Write(buff[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", r.String(), ew.Error())
				os.Exit(1)
			}
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", r.String(), er.Error())
			os.Exit(1)
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(file.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := file.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error: %s\n", flag.Arg(i), err.Error())
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}
