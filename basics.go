package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Rachel")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
