/*
	Although this code is techicaly written by me, 
	it is _heavily_ inspired by the following talk
	by Liz Rice: https://www.youtube.com/watch?v=Utf-A4rODH8
*/

package main 
import "os"

func main() {

	if (len(os.Args) <= 1) {

		panic("Must have an argument"); 
	}

	switch (os.Args[1]) {
		case "run": 
			run() 
	default: 
		panic("Invalid argument.")
	}
}

func run() {}

func must(err error) {

	if err != nil {
		panic(err)
	}
} 