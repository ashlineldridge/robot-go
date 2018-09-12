package main

import (
	"fmt"
	"os"
)

func Check(e error, msg string, a ...interface{}) {
	if e != nil {
		fmt.Fprintln(os.Stderr, msg, a)
		panic(e)
	}
}
