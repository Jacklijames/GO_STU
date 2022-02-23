package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
}

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
}
