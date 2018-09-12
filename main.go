package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var sim = Simulation{Table: Table{Width: 5, Height: 5}}

func main() {
	flag.Parse()
	var scanner *bufio.Scanner
	switch flag.NArg() {
	case 0:
		println("Reading instructions from stdin")
		scanner = bufio.NewScanner(os.Stdin)
	case 1:
		fmt.Printf("Reading instructions from file %v", flag.Arg(0))
		f, err := os.Open(flag.Arg(0))
		defer f.Close()
		Check(err, "Error opening file %v", flag.Arg(0))
		scanner = bufio.NewScanner(bufio.NewReader(f))
	default:
		println("Incorrect usage")
		os.Exit(1)
	}
	for scanner.Scan() {
		cmd := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		ParseCommand(cmd)(&sim)
	}
}
