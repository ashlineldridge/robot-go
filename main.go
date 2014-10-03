package main

import "fmt"

func main() {
	r := Position{x: 1, y: 1, d: North}
	fmt.Println(r)

	p := Position{x: 1, y: 1, d: 3}
	t := Table{width: 5, height: 5}
	fmt.Println(t.IsValid(&p))
}
