package main

import "fmt"

type Position struct {
	x int
	y int
	d Direction
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d,%s", p.x, p.y, p.d)
}
